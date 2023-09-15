﻿using NLog;

namespace Geek.Server.Core.Reference;

/// <summary>
/// 引用池。
/// </summary>
public static partial class ReferencePool
{
    private static readonly NLog.Logger LOGGER = LogManager.GetCurrentClassLogger();
    private static readonly Dictionary<Type, ReferencePool.ReferenceCollection> s_ReferenceCollections = new Dictionary<Type, ReferencePool.ReferenceCollection>();
    private static bool m_EnableStrictCheck = false;

    /// <summary>
    /// 获取或设置是否开启强制检查。
    /// </summary>
    public static bool EnableStrictCheck
    {
        get
        {
            return m_EnableStrictCheck;
        }
        set
        {
            m_EnableStrictCheck = value;
        }
    }

    /// <summary>
    /// 获取引用池的数量。
    /// </summary>
    public static int Count
    {
        get
        {
            return s_ReferenceCollections.Count;
        }
    }

    /// <summary>
    /// 获取所有引用池的信息。
    /// </summary>
    /// <returns>所有引用池的信息。</returns>
    public static ReferencePoolInfo[] GetAllReferencePoolInfos()
    {
        int index = 0;
        ReferencePoolInfo[] results = null;

        lock (s_ReferenceCollections)
        {
            results = new ReferencePoolInfo[s_ReferenceCollections.Count];
            foreach (KeyValuePair<Type, ReferencePool.ReferenceCollection> referenceCollection in s_ReferenceCollections)
            {
                results[index++] = new ReferencePoolInfo(referenceCollection.Key, referenceCollection.Value.UnusedReferenceCount, referenceCollection.Value.UsingReferenceCount, referenceCollection.Value.AcquireReferenceCount, referenceCollection.Value.ReleaseReferenceCount, referenceCollection.Value.AddReferenceCount, referenceCollection.Value.RemoveReferenceCount);
            }
        }

        return results;
    }

    /// <summary>
    /// 清除所有引用池。
    /// </summary>
    public static void ClearAll()
    {
        lock (s_ReferenceCollections)
        {
            foreach (KeyValuePair<Type, ReferencePool.ReferenceCollection> referenceCollection in s_ReferenceCollections)
            {
                referenceCollection.Value.RemoveAll();
            }

            s_ReferenceCollections.Clear();
        }
    }

    /// <summary>
    /// 从引用池获取引用。
    /// </summary>
    /// <typeparam name="T">引用类型。</typeparam>
    /// <returns>引用。</returns>
    public static T Acquire<T>() where T : class, IReference, new()
    {
        return GetReferenceCollection(typeof(T)).Acquire<T>();
    }

    /// <summary>
    /// 从引用池获取引用。
    /// </summary>
    /// <param name="referenceType">引用类型。</param>
    /// <returns>引用。</returns>
    public static IReference Acquire(Type referenceType)
    {
        InternalCheckReferenceType(referenceType);
        return GetReferenceCollection(referenceType).Acquire();
    }

    /// <summary>
    /// 将引用归还引用池。
    /// </summary>
    /// <param name="reference">引用。</param>
    public static void Release(IReference reference)
    {
        if (reference == null)
        {
            LOGGER.Fatal("Reference is invalid.");
            return;
        }

        Type referenceType = reference.GetType();
        InternalCheckReferenceType(referenceType);
        GetReferenceCollection(referenceType).Release(reference);
    }

    /// <summary>
    /// 向引用池中追加指定数量的引用。
    /// </summary>
    /// <typeparam name="T">引用类型。</typeparam>
    /// <param name="count">追加数量。</param>
    public static void Add<T>(int count) where T : class, IReference, new()
    {
        GetReferenceCollection(typeof(T)).Add<T>(count);
    }

    /// <summary>
    /// 向引用池中追加指定数量的引用。
    /// </summary>
    /// <param name="referenceType">引用类型。</param>
    /// <param name="count">追加数量。</param>
    public static void Add(Type referenceType, int count)
    {
        InternalCheckReferenceType(referenceType);
        GetReferenceCollection(referenceType).Add(count);
    }

    /// <summary>
    /// 从引用池中移除指定数量的引用。
    /// </summary>
    /// <typeparam name="T">引用类型。</typeparam>
    /// <param name="count">移除数量。</param>
    public static void Remove<T>(int count) where T : class, IReference
    {
        GetReferenceCollection(typeof(T)).Remove(count);
    }

    /// <summary>
    /// 从引用池中移除指定数量的引用。
    /// </summary>
    /// <param name="referenceType">引用类型。</param>
    /// <param name="count">移除数量。</param>
    public static void Remove(Type referenceType, int count)
    {
        InternalCheckReferenceType(referenceType);
        GetReferenceCollection(referenceType).Remove(count);
    }

    /// <summary>
    /// 从引用池中移除所有的引用。
    /// </summary>
    /// <typeparam name="T">引用类型。</typeparam>
    public static void RemoveAll<T>() where T : class, IReference
    {
        GetReferenceCollection(typeof(T)).RemoveAll();
    }

    /// <summary>
    /// 从引用池中移除所有的引用。
    /// </summary>
    /// <param name="referenceType">引用类型。</param>
    public static void RemoveAll(Type referenceType)
    {
        InternalCheckReferenceType(referenceType);
        GetReferenceCollection(referenceType).RemoveAll();
    }

    private static void InternalCheckReferenceType(Type referenceType)
    {
        if (!m_EnableStrictCheck)
        {
            return;
        }

        if (referenceType == null)
        {
            LOGGER.Fatal("Reference type is invalid.");
            return;
        }

        if (!referenceType.IsClass || referenceType.IsAbstract)
        {
            LOGGER.Fatal("Reference type is not a non-abstract class type.");
            return;
        }

        if (!typeof(IReference).IsAssignableFrom(referenceType))
        {
            LOGGER.Fatal($"Reference type '{referenceType.FullName}' is invalid.");
        }
    }

    private static ReferencePool.ReferenceCollection GetReferenceCollection(Type referenceType)
    {
        if (referenceType == null)
        {
            LOGGER.Fatal("Reference type is invalid.");
        }

        ReferencePool.ReferenceCollection referenceCollection = null;
        lock (s_ReferenceCollections)
        {
            if (!s_ReferenceCollections.TryGetValue(referenceType, out referenceCollection))
            {
                referenceCollection = new ReferencePool.ReferenceCollection(referenceType);
                s_ReferenceCollections.Add(referenceType, referenceCollection);
            }
        }

        return referenceCollection;
    }
}