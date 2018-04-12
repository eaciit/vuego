<template>
    <form style="max-width:350px">
        <div class="row"><div class="input-field col s12">
            <input id="email" type="text" class="validate" placeholder="email" v-model="user.email">
            <label class="active" for="loginID">Email</label>
        </div></div>
        <div class="row">
            <div class="input-field col s12">
                <button class="btn btn-small btn-default col s12" type="button" v-on:click="remindMe">Remind me</button>
            </div>
        </div>
        <div class="row">
            <div class="input-field col s6">
                <a href="#" v-on:click="$emit('change-mode','login')">Login</a>
            </div>
            <div class="input-field col s6 right-align">
                <a href="#" v-on:click="$emit('change-mode','register')">Register</a>
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
    name:"forget",
    data () {
        return {
            user:{
                email:""
            },
            result:{status:"", message:""}
        }
    },
    props:["forgetUrl"],
    methods:{
        remindMe () {
            this.$http.post(this.forgetUrl, this.user).
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
                        this.result.message="Invalid response from " + this.forgetUrl;
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