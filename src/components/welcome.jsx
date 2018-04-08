import {render, Component} from "inferno";

import heartImage from "../images/heart.png";
import "./welcome.sass";


class Welcome extends Component {
    render() {
        return (
            <div id="welcome">
                <div id="heart-container">
                    <img id="heart" src={heartImage} alt="heart" />
                </div>
                <h1>{this.props.children}</h1>
                <h3><a class="text-white" href="https://github.com/UberBosser/GoTemplate">GoTemplate</a></h3>
            </div>
        )
    }
}


export default Welcome;
