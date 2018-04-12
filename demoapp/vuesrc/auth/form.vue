<template>
    <form style="max-width:350px">
        <div class="row"><div class="input-field col s12">
            <input id="loginID" type="text" class="validate" placeholder="your login or email id" v-model="loginData.userID">
            <label class="active" for="loginID">Login ID</label>
        </div></div>

        <div class="row"><div class="input-field col s12">
            <input id="password" type="password" class="validate" placeholder="password" v-model="loginData.password">
            <label class="active" for="password">Password</label>
        </div></div>
    
        <div class="row">
        <div class="switch col s6">
            <label class="active">Remember me</label>
            <label>
            Off
            <input type="checkbox" v-model="loginData.rememberMe">
            <span class="lever"></span>
            On
            </label>
        </div>
        </div>
        <div class="row">
            <div class="input-field col s12">
                <button class="btn btn-small btn-default col s12" type="button" @click="loginDo">Login</button>
            </div>
        </div>
        <div class="row">
            <div class="input-field col s6">
                <a href="#" @click="$emit('change-mode','register')">Register</a>
            </div>
            <div class="input-field col s6 right-align">
                <a href="#" @click="$emit('change-mode','forget')">Forget</a>
            </div>
        </div>
        <div class="row" v-if="result.status!='OK' || result.message!=''">
            <div class="col s12 m5">
                {{ result.message }}
            </div>
        </div>
    </form>
</template>

<script>
import Vue from "vue"
import VueResource from "vue-resource"

Vue.use(VueResource)

export default {
  name: 'loginform',
  props:['loginUrl'],
  data () {
    return {
        loginData: {userID: '', password: '', rememberMe: false},
        result :{status:"",message:""}
    }
  },
  methods: {
      loginDo: function(event){
          this.$http.post(this.loginUrl, this.loginData).
            then(
                r=>{
                    if(r.body==undefined || r.body.Status==undefined){
                        this.result.status="NOK";
                        this.result.message="Invalid response";
                    }else if(r.body.Status!="OK" && r.body.message!=undefined){
                        this.result.status="NOK";
                        this.result.message=r.body.Message;
                    }else if(r.body.Status!="OK"){
                        this.result.status="NOK";
                        this.result.message="Invalid response from " + this.loginUrl + JSON.stringify(r);
                    }else{
                        this.result.status="OK";
                        this.result.message="Login Success"
                    }
                }, 
                e=>{
                    this.result.status="NOK";
                    this.result.message=e;
                });
      }
  }
}
</script>