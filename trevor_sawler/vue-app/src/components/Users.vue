<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">All Users</h1>
      </div>

      <hr />

      <table v-if="ready" class="table table-compact table-striped">
        <thead>
          <tr>
            <th>User</th>
            <th>Email</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="u in this.users" :key="u.id">
            <td>
              <router-link :to="`/admin/users/${u.id}`"
                >{{ u.last_name }}, {{ u.first_name }}</router-link
              >
            </td>
            <td>{{ u.email }}</td>
          </tr>
        </tbody>
      </table>

      <p v-else>Loading...</p>
    </div>
  </div>
</template>

<script>
import Security from '@/components/security'
import notie from 'notie'

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

export default {
  data() {
    return {
      users: [],
      ready: false
    }
  },
  beforeMount() {
    // protected route
    Security.requireToken()

    fetch(`${import.meta.env.VITE_APP_API_URL}/admin/users`, Security.requestOptions(''))
      .then((response) => response.json())
      .then((response) => {
        if (response.error) {
          notie.alert({
            type: 'error',
            text: response.message
          })
        } else {
          sleep(3000).then(() => {
            this.users = response.data.users
            this.ready = true
          })
        }
      })
      .catch((error) => {
        notie.alert({
          type: 'error',
          text: error
        })
      })
  }
}
</script>
