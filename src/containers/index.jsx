import { render, Component } from "inferno";
import css from "../css/global.sass";

import Navbar from "../components/navbar.jsx";


class Page extends Component {
    render() {
        return(
            <div>
                <Navbar /> 
            </div>
        ) 
    }
}

render(<Page />, document.getElementById("app"));
