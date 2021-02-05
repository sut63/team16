import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SearchIcon from '@material-ui/icons/Search'; //search icon
import ArrowBackIcon from '@material-ui/icons/ArrowBack'; // back icon
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
import {
  EntPromotion,
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
  table: {
    minWidth: 1013,
  },
}));

interface promotionsearch {
  promotion: number,
  name: String,
  code: String,
}

const Promotionsearch: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [promotionsearch, setPromotionsearch] = React.useState<Partial<promotionsearch>>({});

  const [promotion, setPromotion] = React.useState<EntPromotion[]>([]);
 
  const getPromotion = async () => {
    const res = await http.listPromotion({});
    setPromotion(res);
  };

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
  });
  
  // Lifecycle Hooks
  useEffect(() => {
    getPromotion();
  }, []);

  // set data to object promotion
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Promotionsearch;
    const { value } = event.target;
    setPromotion({ ...promotion, [name]: value });
    console.log(Promotionsearch);
  };

  // clear input form
  function clear() {
    setPromotionsearch({});
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'NAME':
        alertMessage("error","ชื่อโปรโมชั่น ห้ามมีตัวอีกษรพิเศษ");
        return;
      case 'CODE':
        alertMessage("error","รหัสต้องขึ้นต้น A-Z 2 ตัวและตามด้วย ตัวเลข 3 หลัก");
        return;
      default:
        alertMessage("error","ค้นหาข้อมูลไม่สำเร็จ");
        return;
    }
  }

  function search() {

  }

  // // function search data
  // function search() {
  //   const apiUrl = 'http://localhost:8080/api/v1/promotions';
  //   const requestOptions = {
  //     method: 'GET',
  //     headers: { 'Content-Type': 'application/json' },
  //     body: JSON.stringify(promotion),
  //   };
    
  //   console.log(promotion); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

  //   fetch(apiUrl, requestOptions)
  //     .then(response => response.json())
  //     .then(data => {
  //       console.log(data);
  //       if (data.status === true) {
  //         clear();
  //         Toast.fire({
  //           icon: 'success',
  //           title: 'ค้นหาข้อมูลสำเร็จ',
  //         });
  //       } else {
  //         checkCaseSaveError(data.error.Name)
  //       }
  //     });
  // }

  return (
    <Page theme={pageTheme.home}>

      <Header style={HeaderCustom} title={`ระบบโปรโมชั่น`}>
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
              <div className={classes.paper}>ชื่อโปรโมชั่น</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="กรอกชื่อโปรโมชั่น"
                  name="name"
                  type="string"
                  value={promotionsearch.name} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={12}></Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รหัสโปรโมชั่น</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="กรอกรหัสโปรโมชั่น"
                  name="code"
                  type="string"
                  value={promotionsearch.code} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={12}></Grid>

            <Grid item xs={4}></Grid>
            <Grid item xs={4}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SearchIcon />}
                onClick={search}
              >
                ค้นหา
              </Button>
            </Grid>
            <Grid item xs={4}></Grid>

            <Grid item xs={12}></Grid>

            <Grid item xs></Grid>
            <TableContainer component={Paper}>
              <Table className={classes.table} size="medium" aria-label="a dense table">
                <TableHead>
                  <TableRow>
                    <TableCell>IDs</TableCell>
                    <TableCell align="right">ชื่อโปรโมชั่น</TableCell>
                    <TableCell align="right">เงื่อนไขโปรโมชั่น&nbsp;</TableCell>
                    <TableCell align="right">รหัสโปรโมชั่น&nbsp;</TableCell>
                    <TableCell align="right">วันที่และเวลาหมดอายุ&nbsp;</TableCell>
                    <TableCell align="right">ผู้รับผิดชอบ&nbsp;</TableCell>
                    <TableCell align="right">จำนวนการใช้งานโปรโมชั่น&nbsp;</TableCell>
                    <TableCell align="right">ประเภทปรโมชั่น&nbsp;</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {/* {EntPromotion.map(() => (
                    <TableRow>
                      <TableCell component="th" scope="row">
                        {row.name}
                      </TableCell>
                      <TableCell align="right">{
                        
                      }</TableCell>
                      <TableCell align="right">{row.fat}</TableCell>
                      <TableCell align="right">{row.carbs}</TableCell>
                      <TableCell align="right">{row.protein}</TableCell>
                    </TableRow>
                  ))} */}
                </TableBody>
              </Table>
            </TableContainer>
            <Grid item xs></Grid>

            <Grid item xs={12}></Grid>

            <Grid item xs={4}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<ArrowBackIcon />}
                href="/Promotion"
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

export default Promotionsearch;
