<p align="center">
<img width="128" src="https://p0.meituan.net/csc/f6a87af36769d1ca3ead50a8d711ef9a3459.png" >
</p>
<p align="center">
<a href="https://wails.io/zh-Hans/docs/introduction" target="_blank">Wails文档</a>
<a href="https://vuejs.org/" target="_blank">Vue3</a>
<a href="https://arco.design/vue/docs/start">Arco Design</a>
<a href="https://github.com/run-bigpig/xpvideo/issues" target="_blank">反馈</a>
</p>

### 软件介绍

XPVideo 是一款采用现代化技术栈开发的媒体播放器，它基于 Wails 2.0 框架，集成了 `Arco Design` UI 组件库与 `Vue3` 全家桶 功能简单 适合于个人学习与娱乐。

**重要提醒**: 在开始使用前，请务必详读并同意用户协议，确保遵守相关规定!

<details>
<summary>展开查看用户协议及免责申明</summary>
感谢您选择使用XPVideo(以下简称本软件)，在使用产品和服务之前，请您仔细阅读和理解以下声明:

1. 若您不同意本声明的任何内容，请您立即停止使用本软件。一旦您开始使用本软件产品和服务，则表示您已同意本声明的所有内容。
2. 本软件仅供个人学习、研究和技术交流使用，仅提供展示功能，所有数据资源均由用户自身制作提供，包括但不限于视频网站、媒体分享站点等。本软件无法控制这些资源的合法性、准确性、完整性或可用性，因此不对资源内容的真实性、合法性或适用性负责。
3. 由于数据源为互联网收集制作，我们在此特别提醒, 视频中可能出现的任何第三方广告、产品推广信息等相关内容，均系第三方(含用户)行为植入，非本软件策划或添加。请您在体验过程中保持警惕，对这类信息的真实性及合法性进行自主甄别，如用户遇诈骗因此产生的损失，本平台不承担任何责任。
4. 本软件仅使用Iframe嵌入多家视频平台网站内容，包括但不限于爱奇艺(iqiyi.com)、腾讯视频(v.qq.com)、搜狐视频(tv.sohu.com)、聚力网(pptv.com)、360影视(360kan.com)及芒果TV(mgtv.com)等。对于用户在使用本软件过程中对如上网站进行的任何操作，本软件不承担任何责任。
5. 本软件具备资源嗅探特性，可能会引发第三方数据的隐私和安全风险。用户在使用该特性时，需自行承担可能产生的信息泄露或滥用风险，并对其后果负全部责任。
6. 我们深知您的隐私无价。因此，本软件绝不收集任何用户数据 所有信息均严格本地存储，确保您的数据仅在您掌控之中。此软件不与任何第三方共享您的任何信息。
7. 您在使用本软件时需自行负责所有操作和使用结果。本软件不对您通过使用本软件获取的任何内容负责，包括但不限于媒体资源的准确性、版权合规性、完整性、安全性和可用性。对于任何因使用本软件导致的损失、损害或法律纠纷，不承担任何责任。 8
8. 您在使用本软件时必须遵守您所在国家/地区的相关法律法规，禁止使用本软件进行任何违反法律法规的活动，包括但不限于制作、上传、传播、存储任何违法、侵权、淫秽、诽谤、恶意软件等内容。如您违反相关法律法规，需自行承担法律责任。
9. 本免责声明适用于本软件的所有用户。本软件保留随时修改、更新本声明的权利，并以Github Readme、软件更新等形式通知用户。请您定期查阅并遵守最新的免责声明。

请您在使用本软件之前认真阅读并理解本免责声明的所有内容，感谢您的理解和支持。
</details>

##  法律声明与注意事项

- 强烈倡导合法观影，本软件仅作为播放工具，不涉及资源存储或分发。
- 仅供个人学习交流之用，24小时内请自觉卸载，勿作商业用途。
- 软件提供播放框架，具体播放源需用户自行配置。

## 数据源
 google 搜索 影视采集 然后在对应资源站上去找到MaccmsV10 json接口 配置即可

## 功能亮点
- 没啥亮点,功能简单 适合个人学习

### 软件截图
![列表](https://p0.meituan.net/csc/4bedd6ed84669dd0ae9e962986dd0507617867.jpg)
![播放](https://p0.meituan.net/csc/690050a024ba5e4ee6896bb5fda75275342809.jpg)
![搜索](https://p0.meituan.net/csc/0d38d128b21e7c9f2422abd5a7c9b89198177.jpg)
![设置数据源](https://p0.meituan.net/csc/24bd92c1a1a84f4d3e25a8cf6df6923f58498.jpg)

##  二次开发

```
[1]安装 node.js version18 以上 和 go 1.19 以上
[2]克隆项目  git clone https://github.com/run-bigpig/xpvideo.git
[3]进入项目  cd xpvideo/
[4]安装后端依赖  go mod tidy
[5]进入前端目录  cd frontend/
[6]安装前端依赖  npm install
[7]打开编辑器
[8]修改代码
[9]安装wails-cli  go install github.com/wailsapp/wails/v2/cmd/wails@latest
[10]编译  wails build -clean
[11]编译完成后，进入build/bin目录，找到xpvideo.exe，双击运行即可 (在不同平台上下，编译后的文件名称不同，请根据您的平台选择对应的编译文件)。
[12]如果编译失败，请查看编译日志，根据错误提示进行解决或者查看Wails文档
```

## 贡献

如果您对该项目有任何建议或发现了 Bug，欢迎提交 Issue 或 Pull Request。 请您遵守 BSD 协议，并遵守 Github 的行为规范。

## 许可证
[BSD 3-Clause License](https://opensource.org/licenses/BSD-3-Clause)

## 免责声明
本仓库只为学习研究，如涉及侵犯个人或者团体利益，请与我取得联系，我将主动删除一切相关资料，谢谢！

## 最后
欢迎fork、star、pull request，谢谢！

