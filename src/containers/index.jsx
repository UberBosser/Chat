import {render, Component} from "inferno";

import Socket from "../utils/socket.jsx";

import ErrorAlert from "../components/errorAlert.jsx";
import Chat from "../components/chat.jsx";

import css from "../css/global.sass";


class Page extends Component {
    constructor(props) {
        super(props);
        this.state = {
            connected: false,
            messages: [],
            errors: []
        } 
    }
    componentDidMount() {
        let socket = this.socket = new Socket();
        socket.on("open", this.onOpen.bind(this));
        socket.on("close", this.onClose.bind(this));
        socket.on("receive message", this.receiveMessage.bind(this));
        socket.on("error", this.onError.bind(this));
    }
    onOpen() {
        this.setState({
            connected: true
        });
    }
    onClose() {
        this.setState({
            connected: false
        });
    }
    onError(error) {
        this.setState({
            errors:  [...this.state.errors, error]
        }); 
    }
    sendMessage(message) {
        this.socket.emit("send message", message);
    }
    receiveMessage(message) {
        this.setState({
            messages: [...this.state.messages, message]
        });
    }
    render() {
        return(
            <div class="container-fluid mt-2">
                <ErrorAlert
                    errors={this.state.errors}
                />
                <Chat
                    connected={this.state.connected}
                    messages={this.state.messages}
                    sendMessage={this.sendMessage.bind(this)}
                />
            </div>
        ) 
    }
}

render(<Page />, document.getElementById("app"));
