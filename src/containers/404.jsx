import {render, Component} from "inferno";

import Welcome from "../components/welcome.jsx";

import "../css/global.sass";


class Page extends Component {
    render() {
        return(
            <div>
                <Welcome>404: Page not found.</Welcome>
            </div>
        ) 
    }
}

render(<Page />, document.getElementById("app"));
