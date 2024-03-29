/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/6/26 10:54
* @Description: The file is for
***********************************************************************/

package ledger
//ledger下为区块链账本结构。目前为止单链，考虑增加其扩展性DAG等结构
//在ledger目录下将存放于区块链相关的所有代码及包
//对于比特币系统，每个节点既是交易产生者，也可能是区块制造者
//这意味着节点既是服务的使用者，也是服务的提供者
//但是在这个系统中，似乎是这样的：终端设备享受网关节点提供的身份管理
// 终端向网关产生交易（即身份进出申请），网关间进行决策
//

// 交易池 -> 请求池
// 交易池提取交易打包区块的优先顺序：根据终端设备信用分数决定先后
// 网关节点打包好区块后广播出去