import dva from 'dva'
import { browserHistory } from 'dva/router'
import createLoading from 'dva-loading'

// 1. Initialize
const app = dva({
  history: browserHistory,
  onError (error) {
    app._store.dispatch({type: 'app/loginFail'})
    console.error('app onError -- ', app)
  }
})
// 2. Plugin
app.use(createLoading())
// 3. Model
app.model(require('./models/app'))

// 4. Router
app.router(require('./router'))

// 5. Start
app.start('#root')
