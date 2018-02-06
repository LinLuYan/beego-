import React from 'react';

class UserImg extends React.Component {
	constructor(props, context){
		super(props);
        this.state = {
        	name: "lly",
        	user:[
                {name:"lly"},
                {name:"ly"},
                {name:"y"},
        	]
        }
	}
	componentWillMount(){
		console.log(1)
	}
	componentDidMount(){
		console.log(2)
		this.props.callback()
	}
	componentWillReceiveProps(){
		debugger
	}
	shouldComponentUpdate(){
		debugger
		return true;        // 记得要返回true
	}
	componentWillUpdate(){
		debugger
	}
	componentDidUpdate(){
		debugger
	}
	componentWillUnmount(){
		debugger
	}
	render(){
		return (
            <div>
			    <img alt="图片必须要有这个" src="http://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIqfDSqNHpP5FZpDRT4wjKylcHkibCja3hwqKB0CKVBX3J2cZ7Zvcw45Y2LHIdx43LY8Lh4GDj3xzA/132"/>
			    <p>父组件传递进来的  {this.props.isFrom}</p>
			    <p>
		            <button onClick={this.toParent}>传值给父组件</button>
			    </p>
			    <p>
			        {
			        	this.state.user.map((item,index) => {
                            return (
                                <p key={index}>{index} {item.name}</p>
                            )
			        	})
			        }
			    </p>
			</div>
		)
	}
	toParent(e){
		debugger
	}
}

export default UserImg