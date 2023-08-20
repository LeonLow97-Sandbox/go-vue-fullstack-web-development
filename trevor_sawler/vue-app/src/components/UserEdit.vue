<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">User</h1>
      </div>
    </div>
  </div>
</template>

<script>
import Security from '@/components/security'

export default {
  beforeMount() {
    // protected route
    Security.requireToken()
  }
}
</script>