export default {
	$post(data){
		var xhr = new XMLHttpRequest();

		xhr.open('post', data.url );

		xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");

		xhr.send(JSON.stringify(data.data))

		xhr.onreadystatechange = function(){
            if(xhr.readyState === 4 && xhr.status === 200){
              	data.success(JSON.parse(xhr.response));
            }
            else{
              	data.error();
            }

		}
	}
}