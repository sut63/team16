import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SearchIcon from '@material-ui/icons/Search'; //search icon
import ArrowBackIcon from '@material-ui/icons/ArrowBack'; // back icon
import Divider from '@material-ui/core/Divider';
import Typography from '@material-ui/core/Typography';
import PowerSettingsNewIcon from '@material-ui/icons/PowerSettingsNew'; // log off icon
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
} from '@material-ui/core';
import { DefaultApi } from '../../../api/apis'; // Api Gennerate From Command
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';
//import { Cookies } from '../../Cookie'
import {
  EntMember,
  EntBookcourse,
} from '../../../api/models';

// header css
const HeaderCustom = {
  minHeight: '20px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
  divider: {
    margin: theme.spacing(2, 0),
  },
  table: {
    minWidth: 1013,
  },
}));


const Bookcoursesearch: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  //var ck = new Cookies()
  //var cookieName = ck.GetCookie()


  const [idMember, setidMember] = React.useState<number>(0);
  const [member, setMember] = React.useState<EntMember[]>([]);

  const getMember = async () => {
    const res = await http.listMember({ limit: 10, offset: 0 });
    setMember(res);
  };

  // alert setting
  const [open, setOpen] = React.useState(false);
  const [fail, setFail] = React.useState(false);

//close alert 
const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
  if (reason === 'clickaway') {
      return;
  }
  setFail(false);
  setOpen(false);
};

  // Lifecycle Hooks
  useEffect(() => {
    getMember();
    listBookcourse();
  }, []);

 // CheckIn  
 var lencheckin : number
 const [bookcourse, setBookcourse] = React.useState<EntBookcourse[]>([])
 const getBookcourse = async () => {
     const res = await http.GetBookcoursebyMember({ id:idMember})
     setBookcourse(res)
     lencheckin = res.length
     if (lencheckin > 0){
         setOpen(true)
     }else{
         setFail(true)
     }   
 }
 
 const listBookcourse = async () => {
     const res = await http.listBookcourse({})
     setBookcourse(res)
 }


  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
) => {
    const { value } = event.target;
    setidMember(value);
};

// clear cookies
//function Clears() {
 //   ck.ClearCookie()
 //   window.location.reload(false)
//}
console.log(bookcourse)
// function seach data
function seach() {
  getBookcourse();
}

  return (
    <Page theme={pageTheme.home}>

      <Header style={HeaderCustom} title={`ระบบค้นหาการจองสนาม`}>
        <Grid item xs>
          <Button
            variant="contained"
            color="secondary"
            size="large"
            startIcon={<PowerSettingsNewIcon />}
            href="/"
          >
            sign out
          </Button>
        </Grid>
      </Header>

      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อ User ลูกค้า</div>
            </Grid>

            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือก User Name</InputLabel>
                <Select
                  name="member"
                  value={idMember || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {member.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.mEMBERNAME}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={12}></Grid>

            <Grid item xs={4}></Grid>
            <Grid item xs={4}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SearchIcon />}
                onClick={seach}
              >
                ค้นหา
              </Button>
            </Grid>
            <Grid item xs={4}></Grid>

            <Grid item xs={12}></Grid>

            <Grid item xs={12}>
                    <Divider className={classes.divider} />
                    <Typography variant="subtitle1" gutterBottom>
                        ตาราง Bookcourse:
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">id Bookcourse</TableCell>
                                <TableCell align="center">ชื่อผู้จอง</TableCell>
                                <TableCell align="center">สนามที่จอง</TableCell>
                                <TableCell align="center">เบอร์โทรติดต่อ</TableCell>
                                <TableCell align="center">จำนวณผู้เข้าใช้</TableCell>
                                <TableCell align="center">รายละเอียด</TableCell>
                                <TableCell align="center">เวลาเข้าใช้</TableCell>
                                <TableCell align="center">พนักงานที่กรอกข้อมูล</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {bookcourse.map((item: EntBookcourse) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.member?.mEMBERNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.course?.cOURSE}</TableCell>
                                    <TableCell align="center">{item.pHONE}</TableCell>
                                    <TableCell align="center">{item.aCCESS}</TableCell>
                                    <TableCell align="center">{item.dETAIL}</TableCell>
                                    <TableCell align="center">{item.bOOKTIME}</TableCell>
                                    <TableCell align="center">{item.edges?.employee?.eMPLOYEENAME}</TableCell>
                                </TableRow> 
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
                <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success">
              ค้นหาสำเร็จ
          </Alert>
          </Snackbar>

          <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              ไม่พบข้อมูล
          </Alert>
          </Snackbar>

            <Grid item xs={12}></Grid>

            <Grid item xs={4}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<ArrowBackIcon />}
                href="./Bookcourse"
              >
                ย้อนกลับ
              </Button>
            </Grid>

          </Grid>

        </Container>

      </Content>

    </Page>
  );
};

export default Bookcoursesearch;
