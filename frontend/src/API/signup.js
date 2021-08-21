const signup = async (userCred)=>{
    var data = {
        username : userCred.username,
        password : userCred.password ,
        email: userCred.email 
    };
    var result = await fetch('http://localhost:8080/signup',{
        method :"POST",
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
        })
    return result.json()
}

export default signup