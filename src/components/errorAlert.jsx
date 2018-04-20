import {render, Component} from "inferno";


class ErrorAlert extends Component {
    render() {
        return (
            <div class="container px-0">
                {this.props.errors.map(error => 
                    <div class="alert alert-danger">
                        <strong>
                            Error:&nbsp;
                        </strong>
                        {error}
                    </div>
                )}
            </div>
        ) 
    }
}


export default ErrorAlert;
