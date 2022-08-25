import Vue from 'vue'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
import { message } from 'ant-design-vue'

message.config({
    top: `20px`,
    duration: 2,
    maxCount: 3,
});

Vue.use(Antd)
Vue.prototype.$message = message
