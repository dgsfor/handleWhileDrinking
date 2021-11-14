const app = getApp();
import { MSreApiRequest } from "../../request/index.js"
Page({
  data: {
    checkLogin: false,
    StatusBar: app.globalData.StatusBar,
    CustomBar: app.globalData.CustomBar,
    // SRE 
    sreList:[{
      icon: 'recordfill',
      color: 'red',
      name: '功能1',
      path: ''
    },
    {
      icon: 'discoverfill',
      color: 'purple',
      name: '功能2',
      path: ''
    },
    {
      icon: 'list',
      color: 'purple',
      name: '功能3',
      path: ''
    }],
    // DBA
    dbaList:[{
      icon: 'noticefill',
      color: 'olive',
      name: '功能4'
    },
    {
      icon: 'recordfill',
      color: 'purple',
      name: '功能5',
      path: ''
    }],
    // mtlab
    monitorList:[{
      icon: 'lightauto',
      color: 'red',
      name: '功能6',
      path: ''
    },
    {
      icon: 'delete',
      color: 'olive',
      name: '功能7',
      path: ''
    },
    {
      icon: 'add',
      color: 'olive',
      name: '功能8',
      path: ''
    }],
    gridCol:5,
    skin: false
  },
  onShow: function() {
    this.CheckLogin();
  },

  jumpTo: function(option) {
    var path = option.currentTarget.dataset.path;
    // check login
    if(!this.data.checkLogin){
      wx.showToast({
        title: "登录失效！",
        icon: 'none'
      });
    } else {
      // jump to path
      wx.navigateTo({
        url: '../' + path
      });
    }
    
  },

  async CheckLogin(){
    const res = await MSreApiRequest({url:"/oauth/wx/check_login",method:"GET"})
    if(res.data.code != 401){
      var checkLogin = true
      this.setData({
        checkLogin,
      })
    } else {
      var checkLogin = false
      this.setData({
        checkLogin,
      })
    }
  },
})