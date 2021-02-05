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
  EntEquipmentrental,
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


const Equipmentrentalsearch: FC<{}> = () => {
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
    listEquipmentrental();
  }, []);

 // CheckIn  
 var lencheckin : number
 const [equipmentrental, setEquipmentrental] = React.useState<EntEquipmentrental[]>([])
 const getEquipmentrental = async () => {
     const res = await http.getGetEquipmentrentalbyMember({ id:idMember})
     setEquipmentrental(res)
     lencheckin = res.length
     if (lencheckin > 0){
         setOpen(true)
     }else{
         setFail(true)
     }   
 }
 
 const listEquipmentrental = async () => {
     const res = await http.listEquipmentrental({})
     setEquipmentrental(res)
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
console.log(equipmentrental)
// function seach data
function seach() {
  getEquipmentrental();
}

  return (
    <Page theme={pageTheme.home}>

      <Header style={HeaderCustom} title={`ระบบค้นหาใบยืมอุปกรณ์กีฬา`}>
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

            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.container}>
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
            <Grid item xs>
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
                        ตาราง Equipmentrental:
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">id Equipmentrental</TableCell>
                                <TableCell align="center">ชื่อผู้ยืม</TableCell>
                                <TableCell align="center">ชื่ออุปกรณ์</TableCell>
                                <TableCell align="center">ชนิดอุปกรณ์</TableCell>
                                <TableCell align="center">จำนวนอุปกรณ์</TableCell>
                                <TableCell align="center">วันที่ยืมอุปกรณ์</TableCell>
                                <TableCell align="center">วันที่ครบกำหนดการยืม</TableCell>
                                <TableCell align="center">ชื่อพนักงาน</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {equipmentrental.map((item: EntEquipmentrental) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.member?.mEMBERNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.equipment?.eQUIPMENTNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.equipmenttype?.eQUIPMENTTYPE}</TableCell>
                                    <TableCell align="center">{item.rENTALAMOUNT}</TableCell>
                                    <TableCell align="center">{item.rENTALDATE}</TableCell>
                                    <TableCell align="center">{item.rETURNDATE}</TableCell>
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
                href="./Equipmentrental"
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

export default Equipmentrentalsearch;