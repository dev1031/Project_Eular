import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import { useState } from 'react';
import login from '../API/login';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    display : "flex",
    justifyContent:"center",
    alignItems :"center"
  },
  rootForm: {
      width: '25ch',
      height :"30px",
      padding :"20px",
      marginTop :"10%"
  },
  button: {
    background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
    border: 0,
    width : "100%",
    borderRadius: 3,
    boxShadow: '0 3px 5px 2px rgba(255, 105, 135, .3)',
    color: 'white',
    height: 38,
    padding: '0 30px',
    display : 'flex',
    flexDirection : 'row'
  },
}));

function Login(){
    const classes = useStyles();
    const [username, setUserName] = useState("");
    const [password, setPassword] = useState("");
    const handleSubmit = async (e)=>{
        e.preventDefault();
        var userCred = {
            username , 
            password 
        }
        var response = await login(userCred)
        if (response.InsertedID) {
          window.location.href = "/home"
        }
    }
    return (
      <div className ={classes.root}>
        <form className={classes.rootForm} onSubmit ={handleSubmit}>
          <Grid container spacing={3}>
            <Grid item xs>
              <TextField type ="text" id="outlined-basic" label="UserName" variant="outlined" value = {username} onChange={e => setUserName(e.target.value)}/>
            </Grid>
          </Grid>
          <Grid container spacing={3}>
            <Grid item xs>
              <TextField type = "text" id="outlined-basic" label="Password" variant="outlined" value ={password} onChange = {e =>setPassword(e.target.value)}/>
            </Grid>
          </Grid>
          <Grid container spacing={3}>
            <Grid item xs>
              <Button className={classes.button} value = "Submit" type ="submit">Login</Button>
            </Grid>
          </Grid>
        </form>
      </div>
    );
}

export default Login