import React from 'react';
import { List, InputItem, Button, WhiteSpace } from 'antd-mobile';
import util from '../utils/util'
import styles from './start.css';

class UserStart extends React.Component {
	constructor(props, context){
		super(props);
        this.state = {
        	service:"暂无。。"
        }
	}
	componentWillMount(){
		var that = this;
		util.$post({
			url: "http://" + document.domain + ":9999/v1/user/isLogin",
			success(res){
				if(res.code ==  "0000"){
					console.log("success")
				}
				else{
					that.props.history.push('/')
				}
			},
			error(){
				console.log("isLogin error")
			}
		})
	}
	componentDidMount(){
	}
	render(){
		return (
            <div className={styles['start']}>
                <p className={styles['service']}>服务返回信息: {this.state.service}</p>
                <div className={styles['btn-con']}>
                    <Button onClick={this.start.bind(this)} type="primary" inline>start</Button>
                </div>
			</div>
		)
	}
	start(){
		var that = this;
		util.$post({
			url: "http://" + document.domain + ":9999/v1/user/start",
			success(res){
				if(res.code === "0000"){
					that.setState({
	                	service: res.msg
	                })
				}
				else{
					that.props.history.push('/')
				}
			},
			error(){
				console.log("start error")
			}
		})
	}
}

export default UserStart