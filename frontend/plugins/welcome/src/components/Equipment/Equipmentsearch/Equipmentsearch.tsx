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
  EntEmployee,
  EntEquipment,
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


const Equipmentsearch: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  //var ck = new Cookies()
  //var cookieName = ck.GetCookie()


  const [idEmployee, setidEmployee] = React.useState<number>(0);
  const [employee, setEmployee] = React.useState<EntEmployee[]>([]);

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployee(res);
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
    getEmployee();
    listEquipment();
  }, []);

 // CheckIn  
 var lencheckin : number
 const [equipment, setEquipment] = React.useState<EntEquipment[]>([])
 const getEquipment = async () => {
     const res = await http.getGetEquipmentbyEmployee({ id:idEmployee})
     setEquipment(res)
     lencheckin = res.length
     if (lencheckin > 0){
         setOpen(true)
     }else{
         setFail(true)
     }   
 }
 
 const listEquipment = async () => {
     const res = await http.listEquipment({})
     setEquipment(res)
 }


  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
) => {
    const { value } = event.target;
    setidEmployee(value);
};

// clear cookies
//function Clears() {
 //   ck.ClearCookie()
 //   window.location.reload(false)
//}
console.log(equipment)
// function seach data
function seach() {
  getEquipment();
}

  return (
    <Page theme={pageTheme.home}>

      <Header style={HeaderCustom} title={`ระบบค้นหาอุปกรณ์`}>
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
              <div className={classes.paper}>ชื่ออุปกรณ์</div>
            </Grid>

            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกอุปกรณ์</InputLabel>
                <Select
                  name="employee"
                  value={idEmployee || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {employee.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.eMPLOYEENAME}
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
                        ตาราง Equipment:
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">ID Equipment</TableCell>
                                <TableCell align="center">ชื่อของอุปกรณ์</TableCell>
                                <TableCell align="center">ประเภทของอุปกรณ์</TableCell>
                                <TableCell align="center">จำนวนของอุปกรณ์</TableCell>
                                <TableCell align="center">ลักษณะนามของอุปกรณ์</TableCell>
                                <TableCell align="center">ตำแหน่งการเก็บอุปกรณ์</TableCell>
                                <TableCell align="center">รายละเอียดของอุปกรณ์</TableCell>
                                <TableCell align="center">พนักงานที่กรอกข้อมูล</TableCell>
                                <TableCell align="center">วันที่เพิ่ม</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {equipment.map((item: EntEquipment) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.eQUIPMENTNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.equipmenttype?.eQUIPMENTTYPE}</TableCell>
                                    <TableCell align="center">{item.eQUIPMENTAMOUNT}</TableCell>
                                    <TableCell align="center">{item.edges?.classifier?.eQUIPMENTCLASSIFIER}</TableCell>
                                    <TableCell align="center">{item.edges?.zone?.eQUIPMENTZONE}</TableCell>
                                    <TableCell align="center">{item.eQUIPMENTDETAIL}</TableCell>
                                    <TableCell align="center">{item.edges?.employee?.eMPLOYEENAME}</TableCell>
                                    <TableCell align="center">{item.eQUIPMENTDATE}</TableCell>
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
                href="./Equipment"
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

export default Equipmentsearch;
