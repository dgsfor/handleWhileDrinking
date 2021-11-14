// pages/home/home.js
const app = getApp()
var that = this
Page({

  /**
   * 页面的初始数据
   */
  data: {
    // 轮播图
    swiperList:[
      {
        "image_desc":"K8S",
        "image_url":"https://whyme.obs.cn-north-4.myhuaweicloud.com/miniapps-lunbo/lunbo1.jpeg"
      },
      {
        "image_desc":"Devops",
        "image_url":"https://whyme.obs.cn-north-4.myhuaweicloud.com/miniapps-lunbo/lunbo2.jpg"
      },
      {
        "image_desc":"meitu",
        "image_url":"https://whyme.obs.cn-north-4.myhuaweicloud.com/miniapps-lunbo/lunbo3.jpeg"
      },
      {
        "image_desc":"ads",
        "image_url":"https://whyme.obs.cn-north-4.myhuaweicloud.com/miniapps-lunbo/lunbo4.jpeg"
      },
    ],
    // 运维产品列表
    OpsProductList: [
      {
        "product_name":"k8s",
        "product_icon_url":""
      },
      {
        "product_name":"jenkins",
        "product_icon_url":""
      },
      {
        "product_name":"gitlabci",
        "product_icon_url":""
      },
    ],
    // 公司产品列表
    productList: [
      {
        "product_name":"美图秀秀",
        "product_icon_url":"https://corp-static.meitu.com/corp-new/6f5c5e561f32eaacb69d17d5e84e12e1_1586784444.png"
      },
      {
        "product_name":"美颜相机",
        "product_icon_url":"https://corp-static.meitu.com/corp-new/96b070740003f5f1f629263f7cdfed56_1586784428.png"
      },
      {
        "product_name":"美拍",
        "product_icon_url":"https://corp-static.meitu.com/corp-new/abe66f25a9d4e35f520b47029be44e16_1586784414.png"
      },
      {
        "product_name":"美妆相机",
        "product_icon_url":"https://corp-static.meitu.com/corp-new/76d3c6f28df3529946ccb64396f9b1ef_1586784368.png"
      },
      {
        "product_name":"BeautyPlus",
        "product_icon_url":"https://corp-static.meitu.com/corp-new/09a0953abc3ae54bb88331700f169ad4_1586784215.png"
      },
      {
        "product_name":"AirBrush",
        "product_icon_url":"https://corp-static.meitu.com/corp-new/f3200091c4740c9bc3bc1f688de4378d_1586784198.png"
      },
      {
        "product_name":"VCUS",
        "product_icon_url":"https://corp-static.meitu.com/corp-new/925fa3c698f7cb6af45a234c780c49b9_1586784181.png"
      },
      {
        "product_name":"Pomelo",
        "product_icon_url":"https://corp-static.meitu.com/corp-new/35d17273d0dcd6b713254cf1e5965025_1586784156.png"
      },
      {
        "product_name":"AirVid",
        "product_icon_url":"https://corp-static.meitu.com/corp-new/4ae22e5c9b923e5ed3c3cf1b15f56111_1586784080.png"
      },
    ],
    colorArr: [],
    colorCount: 1

  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function(options) {
    that = this
    var colors = app.globalData.ColorList;
    that.setData({
      colorArr: colors,
      colorCount: colors.length
    })
  },
})