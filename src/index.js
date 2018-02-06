import React from 'react';
import ReactDOM from 'react-dom';
import {
  HashRouter as Router,
  Route,
  Link,
  Switch
} from 'react-router-dom';

import './index.css'
import Login from './components/login.js'
import Start from './components/start.js'

let content = {
	position:'absolute',
	top:0,
	left:0,
	width:"100%",
	height:"100%"
}

const BasicExample = () => (
	<div style={content}>
        <Router>
	        <Switch>
		        <Route exact path="/" component={Login} />
		        <Route path="/start" component={Start} />
	        </Switch>
	    </Router>
    </div>
)

ReactDOM.render(<BasicExample/>, document.getElementById('root'));