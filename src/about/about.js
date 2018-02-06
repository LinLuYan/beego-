import React from 'react';
import UserImg from './components/userImg'

const user = {
	name:'lly',
	age:26
}

function formatName(user){
	return user.name + "," + user.age;
}

//以回调形式，子组件传值给父组件
function notify(){
	alert("从子传递到父成功")
}


class About extends React.Component{
	render(){
		return(
			<div>
			    <h2>About</h2>
			    <p>{formatName(user)}</p>
			    <UserImg isFrom="是的" callback={notify} />
			</div>
		)
	}
	
}

export default About