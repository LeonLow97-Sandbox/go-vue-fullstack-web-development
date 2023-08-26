# Vue

## Making **JWT Token** a global variable

```js
// store.js to store the value of the token
import { reactive } from 'vue';

export const store = reactive({
  token: '',
});
```

## Navigation Guards

- Navigation guards provided by Vue router are primarily used to guard navigation either by redirecting it or canceling it.
- Protect the routes when user is unauthorized.

```js
const router = createRouter({ ... })

router.beforeEach((to, from) => {
  // ...
  // explicitly return false to cancel the navigation
  return false
})
```

## Transition

- [TransitionGroup](https://vuejs.org/guide/built-ins/transition-group.html)
- [Transition](https://vuejs.org/guide/built-ins/transition.html)

# Golang

## How Go works with JSON Files

---

#### Useful Links

- [Go Playground MarshalIndent](https://go.dev/play/p/2SP7H3ybk9D)
- [Go Playground Unmarshal](https://go.dev/play/p/NaMNv6bzEXA)
- [GoByExample JSON](https://gobyexample.com/json)
- [Go Default Encoding Types (Go Type - JSON Type)](https://blog.boot.dev/golang/json-golang/#default-type)

---

#### Go Default Encoding Types

JSON and GO Types don't match up 1 to 1. Below is a table that describes the type relationships when encoding and decoding.

|   Go Type   | JSON Type |
| :---------: | :-------: |
|    bool     |  boolean  |
|   float64   |  number   |
|   string    |  string   |
| nil pointer |   null    |

---

## Error Codes in PostgreSQL

- [Error Codes in Postgres](https://www.postgresql.org/docs/current/errcodes-appendix.html)
