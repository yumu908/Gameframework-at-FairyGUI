﻿using FairyGUI;
using GameFramework.UI;
using UnityGameFramework.Runtime;

namespace GameLogic.Login
{
    public class LoginForm : UIForm<UI_login_main>
    {
        protected override UIFormConfig CreateConfig()
        {
            return UIFormConfig.Create(UI_login_main.PKG_NAME, UI_login_main.RES_NAME,  groupEnum:UIGroupEnum.PANEL);
        }

        protected override void OnOpen()
        {
            Timers.inst.Add(3f,1, _ => { View.m_title.text = ConfigLoader.Instance.Tables.TbItem.DataList[0].Name;});
        }
    }
}