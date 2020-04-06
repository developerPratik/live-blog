import React from 'react';
import './App.css';

interface IState {
  data: any;
}


class App extends React.Component<{}, IState> {
  ws: (WebSocket | null) = null;

  state = {
    data: ""
  }

  // constructor(props: any) {
  //   super(props);
  // }

  componentDidMount() {
    this.ws = new WebSocket("ws://localhost:8000/ws");
    this.ws.onopen = this.sendData;
    this.ws.onmessage = this.handleMessages;
  }

  sendData = () => {
  }



  handleMessages = (data: any) => {
    console.log("new message", data);
    this.setState({
      data: data.data
    })
  }

  render() {

    return <div>{this.state.data}</div>
  }

  componentWillUnmount() {
    if (this.ws) {
      this.ws.close()
    }
  }


}


export default App;
