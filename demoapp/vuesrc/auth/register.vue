<template>
       <form style="max-width:350px">
        <div class="row"><div class="input-field col s12">
            <input id="loginID" type="text" class="validate" placeholder="login id" v-model="user.userID">
            <label class="active" for="loginID">Login ID</label>
        </div></div>

        <div class="row"><div class="input-field col s12">
            <input id="email" type="text" class="validate" placeholder="email" v-model="user.email">
            <label class="active" for="loginID">Email</label>
        </div></div>

        <div class="row"><div class="input-field col s12">
            <input id="password" type="password" class="validate" placeholder="password" v-model="user.password">
            <label class="active" for="password">Password</label>
        </div></div>

        <div class="row"><div class="input-field col s12">
            <input id="passwordConfirm" type="password" class="validate" placeholder="confirm password" v-model="user.confirm_password">
            <label class="active" for="passwordConfirm">Confirm Password</label>
        </div></div>
    
          <div class="row">
            <div class="input-field col s12">
                <button class="btn btn-small btn-default col s12" type="button" v-on:click="registerDo">Register</button>
            </div>
        </div>
        <div class="row">
            <div class="input-field col s6">
                <a href="#" v-on:click="$emit('change-mode','login')">Login</a>
            </div>
            <div class="input-field col s6 right-align">
                <a href="#" v-on:click="$emit('change-mode','forget')">Forget</a>
            </div>
        </div>
        <div class="row" v-if="result.status!='OK' && result.message!=''">
            <div class="col s12 m5">
                {{ result.message }}
            </div>
        </div>
    </form>
</template>

<script>

export default{
    name:"register",
    data () {
        return {
            user:{userID:"", password:"", passwordConfirm:"", email:""},
            result:{status:"", message:""}
        }
    },
    props: ["registerUrl"],
    methods:{
        registerDo : function(){
            this.$http.post(this.registerUrl, this.user).
            then(
                r=>{
                    if(r==undefined || r.status==undefined){
                        this.result.status="NOK";
                        this.result.message="Invalid response";
                    }else if(r.status!="OK" && r.message!=undefined){
                        this.result.status="NOK";
                        this.result.message=r.message;
                    }else if(r.status!="OK"){
                        this.result.status="NOK";
                        this.result.message="Invalid response from " + this.registerUrl;
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