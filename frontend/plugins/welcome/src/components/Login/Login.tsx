import React, { FC, useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Paper from '@material-ui/core/Paper';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Swal from 'sweetalert2'; // alert
import { DefaultApi } from '../../api/apis';
import { EntEmployee } from '../../api';
import { Cookies } from '../../cookies';
import { Link as RouterLink } from 'react-router-dom';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © ระบบสถานกีฬา | '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  root: {
    height: '98vh',
  },
  image: {
    backgroundImage: 'url(https://www.nakornchiangmaigames.com/files/com_stadium/2018-10_92f75a600ad84f3.jpg)',
    backgroundRepeat: 'no-repeat',
    backgroundColor:
      theme.palette.type === 'light' ? theme.palette.grey[50] : theme.palette.grey[900],
    backgroundSize: 'cover',
    backgroundPosition: 'center',
  },
  paper: {
    margin: theme.spacing(8, 4),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

interface Login {
  username: string;
  password: string;
}

const Login: FC<{}> = () => {
  const http = new DefaultApi();
  const classes = useStyles();

  const [Name, setName] = React.useState(String);
  const [Password, setPassword] = React.useState(String);
  const [login, setLogin] = React.useState<Partial<Login>>({});
  const [employee, setEmployee] = React.useState<EntEmployee[]>([]);

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 20, offset: 0 });
    setEmployee(res);
  };

  //console.log(employee);
  useEffect(() => {
    getEmployee();
    
  }, []);

  var ck = new Cookies();
  var check : boolean
  const [path, setPath] = React.useState("");

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });
  console.log(path);
  
  function redirecLogin() {
    check = ck.CheckLogin(employee,Name,Password)
    console.log("check => "+check)
    if (check === true) {
      Toast.fire({
        icon: 'success',
        title: 'เข้าสู่ระบบสำเร็จ',
      });
      setPath("/WelcomePage")
      ck.SetCookie("user_email",Name,30)
      ck.SetCookie("user_id",ck.SetID(employee,Name,Password),30)
      ck.SetCookie("user_role","employee",30)
      window.location.reload(false)
    }else if(check === false){
      Toast.fire({
        icon: 'error',
        title: 'เข้าสู่ระบบไม่สำเร็จ',
      });
      setPath("/")
    }
  }

  const NameChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setName(event.target.value as string);
    console.log(Name);
  };
  const PasswordChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPassword(event.target.value as string);
  };

  return (
    <Grid container component="main" className={classes.root}>
      <CssBaseline />
      <Grid item xs={false} sm={4} md={7} className={classes.image} />
      <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
        <div className={classes.paper}>
          <Avatar className={classes.avatar}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            LOGIN 
          </Typography>
          <form className={classes.form} noValidate>
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              id="username"
              label="Username"
              name="username"
              autoComplete="username"
              autoFocus
              value={Name || ''}
              onChange={NameChange}

            />
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
              value={Password || ''}
              onChange={PasswordChange}

            />
            <FormControlLabel
              control={<Checkbox value="remember" color="primary" />}
              label="Remember me"
            />
            <Button
              fullWidth
              variant="contained"
              color="primary"
              type="submit"
              className={classes.submit}
               onClick={redirecLogin}
              component={RouterLink}
              to={path}
            >
              Sign In
            </Button>
            <Grid container>
             
            </Grid>
            <Box mt={5}>
              <Copyright />
            </Box>
          </form>
        </div>
      </Grid>
    </Grid>
  );
};

export default Login;