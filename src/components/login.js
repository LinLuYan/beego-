import React from 'react';
import { List, InputItem, Button, WhiteSpace, Toast } from 'antd-mobile';



import crypto from 'crypto'
import util from '../utils/util'
import styles from './login.css';
import logo from './logo.svg';

class UserLogin extends React.Component {
	constructor(props, context){
		super(props);
        this.state = {
        	account:"",
        	passWord:""
        }
	}
	componentWillMount(){
		console.log(this)
	}
	componentDidMount(){
		console.log(2)
	}
	render(){
		return (
            <div className={styles['login']}>
                <div>
                    <div className={styles['title']}>
                        <img src={logo} className={styles['App-logo']} alt="logo" />
                    </div>
                    <List>
                        <InputItem onChange={this.nameInputHandler.bind(this)}>账号</InputItem>
                        <InputItem type="passWord" onChange={this.passWordInputHandler.bind(this)}>密码</InputItem>
                    </List>
                    <div className={styles['btn-con']}>
                        <WhiteSpace />
                        <WhiteSpace />
                        <Button onClick={this.submitHandler.bind(this)} type="primary" inline>登录</Button>  
                    </div>
                </div>
			</div>
		)
	}
	nameInputHandler(val){
		this.setState({
			account: val
		})
	}
	passWordInputHandler(val){
		this.setState({
			passWord: val
		})
	}
	submitHandler(e){
		var that = this;
		util.$post({
			url: "http://" + document.domain + ":9999/v1/user/login",
			data: {
				account: that.md5Encypt(that.state.account),
				passWord: that.md5Encypt(that.state.passWord)
			},
			success(res){
                if(res.code === "0000" || res.code === "0001"){
                	that.props.history.push('/start')
                }
				else{
                    Toast.fail(res.msg)					
				}
			},
			error(){
				console.log("error")
			}
		})
	}
	md5Encypt(string){
        var md5 = crypto.createHash("md5");
        md5.update(string);
        var stringMd5 = md5.digest("hex")
        return stringMd5
    }
}

export default UserLogin