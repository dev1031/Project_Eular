import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';


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

function SignUp(){
    const classes = useStyles();
    return (
      <div className ={classes.root}>
        <form className={classes.rootForm} noValidate autoComplete="off">
          <Grid container spacing={3}>
            <Grid item xs>
              <TextField id="outlined-basic" label="UserName" variant="outlined" />
            </Grid>
          </Grid>
          <Grid container spacing={3}>
            <Grid item xs>
              <TextField id="outlined-basic" label="Password" variant="outlined" />
            </Grid>
          </Grid>
          <Grid container spacing={3}>
            <Grid item xs>
              <TextField id="outlined-basic" label="Email" variant="outlined" />
            </Grid>
          </Grid>
          <Grid container spacing={3}>
            <Grid item xs>
              <Button className={classes.button}>SIGNUP</Button>
            </Grid>
          </Grid>
        </form>
      </div>
    );
}

export default SignUp