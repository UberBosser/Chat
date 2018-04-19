import {render, Component, linkEvent} from "inferno";

import Message from "./message.jsx";


class Chat extends Component {
    constructor(props) {
        super(props);
        this.state = {
            nameInputValue: "",
            messageInputValue: ""
        };
    }
    handleClick(props, event) {
        if (this.state.messageInputValue != "") {
            if (this.state.nameInputValue == "") {
                this.props.sendMessage({name: "Guest", message: this.state.messageInputValue});
                this.setState({
                    messageInputValue: ""
                });
            } else {
                this.props.sendMessage({name: this.state.nameInputValue, message: this.state.messageInputValue});
                this.setState({
                    messageInputValue: ""
                });
            }
        }
    }
    updateName(event) {
        this.setState({
            nameInputValue: event.target.value 
        });
    }
    updateMessage(event) {
        this.setState({
            messageInputValue: event.target.value 
        });
    }
    render() {
        return (
            <div class="container p-0">
                <div class="container">
                    <div class="bg-white rounded border p-2 mb-1 align-middle row justify-content-between">
                        <div class="font-weight-bold d-inline">Chat</div>
                        {this.props.connected ? <div class="d-inline text-success">Connected</div> : <div class="d-inline text-danger">Disconnected</div>}
                    </div>
                </div>
                {this.props.messages.map(message => 
                    <Message {...message} />
                )}
                <div class="container">
                    <div class="row">
                        <input class="form-control col-2" type="text" placeholder="Guest" value={this.state.nameInputValue} onInput={this.updateName.bind(this)}/>
                        <input class="form-control col-8" type="text" placeholder="Message" value={this.state.messageInputValue} onInput={this.updateMessage.bind(this)}/>
                        <button class="btn btn-primary col-2" type="button" onClick={this.handleClick.bind(this)}>Submit</button>
                    </div>
                </div>
            </div> 
        ) 
    }
}


export default Chat;
