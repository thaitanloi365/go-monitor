import {message} from 'antd'

export default {
  onError(e, a) {
    e.preventDefault()
    console.log(e)
    if (e.message) {
      message.error(e.message)
    }
    else {
      /* eslint-disable */
      console.error(e)
    }
  },
}
