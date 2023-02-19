import React from 'react'

import { LoadConfig } from '../../wailsjs/go/api/App'

export default
class Left extends React.Component {
  async componentDidMount() {
    const res = await LoadConfig()
    if(res.Err) {
      console.error(res)
    }
  }
  render() {
    return <div>left</div>
  }
}