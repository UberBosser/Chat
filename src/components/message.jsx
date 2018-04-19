import {render, Component} from "inferno";


class Message extends Component {
    render() {
        return (
            <div class="container border rounded bg-white p-2 mb-1">
                <div class="align-middle">
                    <span class="font-weight-bold">{this.props.name + " : "}</span>
                    {this.props.message}
                </div>
            </div>
        ) 
    }
}


export default Message;
