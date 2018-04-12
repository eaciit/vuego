<template>
<div id="login">
    <auth login-url="/acl/auth" register-url="/acl/register" forget-url="/acl/forget" />
</div>
</template>

<script>
import auth from './auth/auth.vue'
export default {
  name: 'login',
  components:{
      "auth":auth
  },
  data () {
    return {}
  }
}
</script>
