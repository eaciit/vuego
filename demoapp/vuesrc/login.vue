<template>
<div id="login">
    <auth login-url="/acl/auth" register-url="/acl/register" forget-url="/acl/forget" sourceUrl="/" />
</div>
</template>

<script>
import auth from './auth/auth.vue'
export default {
  name: 'login',
  components:{
      "auth":auth
  },
  methods:{
  },
  data () {
    return {}
  }
}
</script>

<style>
  #login{
     margin:0px auto;
  }
</style>