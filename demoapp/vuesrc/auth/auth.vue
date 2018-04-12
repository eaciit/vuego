<template>
    <div>
        <register @change-mode='onModeChanged' v-if="mode=='register'" :register-url='registerUrl' />
        <login @change-mode='onModeChanged' v-if="mode=='login'" :login-url='loginUrl' />
        <forget @change-mode='onModeChanged' v-if="mode=='forget'" :forget-url='forgetUrl'  />
    </div>
</template>

<script>
import Vue from 'vue'
import VueRes from 'vue-resource'
import RegisterForm from './register.vue'
import LoginForm from './form.vue'
import ForgetForm from './forget.vue'

Vue.use(VueRes)

export default {
    name:"auth",
    data(){
        return {
            mode:"login"
        }
    },
    props:['loginUrl','forgetUrl','registerUrl'],
    components:{
        "login":LoginForm,
        "register":RegisterForm,
        "forget":ForgetForm
    },
    methods:{
        onModeChanged:function(mode){
            this.mode=mode;
        }
    }
}
</script>