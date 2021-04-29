<template>
  <div class="admin-container">
    <div class="input-wrap">
      <div class="btn-wrap">
        <!-- 充 值 -->
        <el-button type="primary" @click="handleRecharge">Recharge</el-button>
        <!-- 转 账 -->
        <el-button type="primary" @click="handleTransfer">Transfer</el-button>
        <!-- 调用智能合约的接口，存储字符串数据 -->
        <el-tooltip class="item" effect="dark" content="Call the interface of the smart contract and store the string data" placement="top-start">
          <el-button type="primary" @click="dialogVisibleSetContracts = true">SET</el-button>
        </el-tooltip>
        <!-- 调用智能合约的接口，获取字符串数据 -->
        <el-tooltip class="item" effect="dark" content="Call the interface of the smart contract and get the string data" placement="top-start">
          <el-button type="primary" @click="getContracts">GET</el-button>
        </el-tooltip>
      </div>
      <div class="search-wrap" />
    </div>
    <el-form :inline="true" :model="list" class="demo-form-inline">
      <!-- 区块链地址 -->
      <el-form-item label="Blockchain Address">
        <el-input v-model="list.addr" class="addr-wrap" disabled />
      </el-form-item>
      <!-- 账户余额 -->
      <el-form-item label="Account Balance">
        <el-input v-model="list.value" disabled />
      </el-form-item>
    </el-form>
    <el-table
      v-loading="listLoading"
      :data="list.tx_list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <!-- 序号-Serial number -->
      <el-table-column label="Serial number" width="95" align="center">
        <template slot-scope="scope">{{ scope.$index }}</template>
      </el-table-column>
      <el-table-column label="Sender" width="120" align="center">
        <template slot-scope="scope">{{ scope.row.SendName }}</template>
      </el-table-column>
      <el-table-column label="Receiver" width="120" align="center">
        <template slot-scope="scope">{{ scope.row.ReciveName }}</template>
      </el-table-column>
      <el-table-column label="Transaction Type" width="120" align="center">
        <template slot-scope="scope">{{ scope.row.Typa }}</template>
      </el-table-column>
      <el-table-column label="Amount" align="center">
        <template slot-scope="scope">{{ scope.row.Value }}</template>
      </el-table-column>
      <el-table-column label="Trading HASH" align="center">
        <template slot-scope="scope">{{ scope.row.Hash }}</template>
      </el-table-column>
      <!-- 操作-Operation -->
      <el-table-column
        class-name="status-col"
        label="Operation"
        width="150"
        align="center"
      >
        <template slot-scope="scope">
          <!-- 查 看-Check See -->
          <el-button
            type="text"
            size="mini"
            @click="handleShow(scope.row)"
          >Check See</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 交易详情-Transaction details -->
    <el-dialog title="交易详情" :visible.sync="dialogVisibleDetail">
      <el-form ref="detail" :model="detail" label-width="120px">
        <!-- 交易HASH-Trading HASH -->
        <el-form-item label="Trading HASH" prop="Hash">
          <el-input
            v-model="detail.Hash"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 发送者-Sender -->
        <el-form-item label="Sender" prop="SendName">
          <el-input
            v-model="detail.SendName"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 接收者-Receiver -->
        <el-form-item label="Receiver" prop="ReciveName">
          <el-input
            v-model="detail.ReciveName"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 金额 -->
        <el-form-item label="Amount" prop="Value">
          <el-input
            v-model="detail.Value"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 交易类型-Transaction Type -->
        <el-form-item label="Transaction Type" prop="Typa">
          <el-input
            v-model="detail.Typa"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 发送地址--Send Address -->
        <el-form-item label="Send Address" prop="SendAddr">
          <el-input
            v-model="detail.SendAddr"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 接收地址-Receiving Address -->
        <el-form-item label="Receiving Address" prop="ReciveAddr">
          <el-input
            v-model="detail.ReciveAddr"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 交易发起时间-Transaction Initiation Time -->
        <el-form-item label="Transaction Initiation Time" prop="CreatedAt">
          <el-input
            v-model="detail.CreatedAt"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 交易更新时间-Trading Update Time -->
        <el-form-item label="Trading Update Time" prop="UpdatedAt">
          <el-input
            v-model="detail.UpdatedAt"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="Trading Update Time" prop="UpdatedAt">
          <el-input
            v-model="detail.UpdatedAt"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 交易区块高度-Trading Block Height -->
        <el-form-item label="Trading Block Height" prop="block_height">
          <el-input
            v-model="detail.block_height"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 时间戳-Timestamp -->
        <el-form-item label="Timestamp" prop="timestamp">
          <el-input
            v-model="detail.timestamp"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- 随机值-Random values -->
        <el-form-item label="随机值" prop="nonce">
          <el-input
            v-model="detail.nonce"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- GAS价格-GAS Price -->
        <el-form-item label="GAS价格" prop="gas_price">
          <el-input
            v-model="detail.gas_price"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
        <!-- GAS数量-Number of GAS -->
        <el-form-item label="Number of GAS" prop="gas">
          <el-input
            v-model="detail.gas"
            disabled
            type="text"
            autocomplete="off"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogClose">Close</el-button>
      </div>
    </el-dialog>

    <!-- 转账信息-Transfer Information -->
    <el-dialog title="Transfer Information" :visible.sync="dialogVisibleTransfer">
      <el-form ref="transferFrom" :model="transferFrom" label-width="120px">
        <!-- 转账对象-Transfer recipients -->
        <el-form-item label="Transfer recipients" prop="to">
          <el-input v-model="transferFrom.to" type="text" autocomplete="off" />
        </el-form-item>
        <!-- 转账金额-Transfer Amount -->
        <el-form-item label="Transfer Amount" prop="value">
          <el-input v-model="transferFrom.value" type="text" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" plain size="medium" @click="handleTransferCore">Transfer</el-button>
        <el-button @click="dialogClose">Cancel</el-button>
      </div>
    </el-dialog>

    <!-- 存储字符串信息-Storing String Information -->
    <el-dialog title="Storing String Information" :visible.sync="dialogVisibleSetContracts">
      <el-form ref="setContractsForm" :model="setContractsForm" label-width="120px">
        <el-form-item label="String Information" prop="value">
          <el-input v-model="setContractsForm.value" type="text" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <!-- 存 储-Storage -->
        <el-button type="primary" plain size="medium" @click="setContracts">Storage</el-button>
        <!-- 取 消-Cancel -->
        <el-button @click="dialogClose">Cancel</el-button>
      </div>
    </el-dialog>

    <!-- 获取的字符串信息-Acquired string information -->
    <el-dialog title="Acquired string information" :visible.sync="dialogVisibleGetContracts">
      <el-form ref="getContractsForm" :model="getContractsForm" label-width="120px">
        <!-- <el-form-item label="字符串信息" prop="value"> -->
        <el-form-item label="String Information" prop="value">
          <el-input v-model="getContractsForm.value" type="text" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <!-- <el-button @click="dialogClose">关 闭</el-button> -->
        <el-button @click="dialogClose">Close</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  getInfo,
  recharge,
  transfer,
  setContracts,
  getContracts,
  gettx
} from 'api/user'
import { getToken } from 'utils/auth'

export default {
  data() {
    const token = getToken()
    return {
      list: {},
      search: '',
      listLoading: false,
      dialogVisibleDetail: false,
      dialogVisibleTransfer: false,
      dialogVisibleSetContracts: false,
      dialogVisibleGetContracts: false,
      detail: {},
      transferFrom: {
        to: '',
        value: '',
        token: token
      },
      setContractsForm: {
        value: '',
        token: token
      },
      getContractsForm: {
        value: ''
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    // 获取个人信息
    fetchData() {
      this.listLoading = true
      getInfo({
        token: getToken()
      })
        .then(response => {
          this.list = response.data || {}
          this.listLoading = false
        })
        .catch(error => {
          this.listLoading = false
          console.log(error)
          this.$message({
            message: 'Failed to get!', // '获取失败!',
            type: 'error'
          })
        })
    },
    // 处理查看详情
    handleShow(item) {
      gettx({ token: getToken(), hash: item.Hash }).then(response => {
        const data = response.data.tx
        this.detail = item
        console.log(data.timestamp)
        this.detail.block_height = data.block_height
        this.detail.timestamp = data.timestamp
        this.detail.nonce = data.nonce
        this.detail.gas_price = data.gas_price
        this.detail.gas = data.gas
        this.dialogVisibleDetail = true
      }).catch(error => {
        console.log(error)
        this.dialogVisibleDetail = true
        this.$message({
          message: 'Block query failure!', // '区块查询失败!',
          type: 'error'
        })
      })
    },
    // 处理充值
    handleRecharge() {
      recharge({ token: getToken() })
        .then(response => {
          this.listLoading = false
          this.fetchData()
          this.$message({
            message: 'Successful recharge!', // '充值成功!',
            type: 'success'
          })
        })
        .catch(error => {
          this.listLoading = false
          console.log(error)
          this.$message({
            message: 'Failed to recharge!', // '充值失败!',
            type: 'error'
          })
        })
    },
    // 处理转账弹窗
    handleTransfer() {
      this.dialogVisibleTransfer = true
    },
    // 处理转账
    handleTransferCore() {
      transfer(this.transferFrom)
        .then(response => {
          this.listLoading = false
          this.fetchData()
          this.dialogClose()
          this.$message({
            message: 'Transfer successful', // '转账成功!',
            type: 'success'
          })
        })
        .catch(error => {
          this.listLoading = false
          console.log(error)
          this.$message({
            message: 'Transfer failed!', // '转账失败!',
            type: 'error'
          })
        })
    },
    // 存入合约
    setContracts() {
      setContracts(this.setContractsForm)
        .then(response => {
          this.dialogVisibleSetContracts = false
          this.fetchData()
          this.$message({
            message: 'Deposit successful!', // '存入成功!',
            type: 'success'
          })
        })
        .catch(error => {
          this.dialogVisibleSetContracts = false
          console.log(error)
          this.$message({
            message: 'Deposit failed!', // '存入失败!',
            type: 'error'
          })
        })
    },
    // 从合约中获取
    getContracts() {
      getContracts({ token: getToken() })
        .then(response => {
          this.dialogVisibleGetContracts = true
          this.getContractsForm.value = response.data.value
        })
        .catch(error => {
          this.dialogVisibleGetContracts = false
          console.log(error)
          this.$message({
            message: 'Failed to get!', // '获取失败!',
            type: 'error'
          })
        })
    },
    // 处理交易详情\转账信息对话框关闭
    dialogClose() {
      this.dialogVisibleDetail = false
      this.dialogVisibleTransfer = false
      this.dialogVisibleGetContracts = false
      this.transferFrom = {
        to: '',
        value: '',
        token: getToken()
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.admin {
  &-container {
    margin: 30px;
    .addr-wrap {
      width: 400px;
    }
    .input-wrap {
      display: flex;
      justify-content: space-between;
      align-content: center;
      padding-bottom: 1vw;
      .input-item {
        width: 15vw;
      }
    }
  }
}

</style>
