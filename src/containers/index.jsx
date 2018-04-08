import {render, Component} from "inferno";

import Welcome from "../components/welcome.jsx";

import css from "../css/global.sass";


class Page extends Component {
    render() {
        return(
            <div>
                <Welcome>You are all set!</Welcome>
            </div>
        ) 
    }
}

render(<Page />, document.getElementById("app"));
