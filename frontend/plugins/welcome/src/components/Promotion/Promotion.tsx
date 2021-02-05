import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import SearchIcon from '@material-ui/icons/Search'; //search icon
import PowerSettingsNewIcon from '@material-ui/icons/PowerSettingsNew'; //log off icon
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
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import {
  EntPromotiontype,
  EntPromotionamount,
  EntEmployee,
} from '../../api/models';

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
}));

interface promotion {
  promotiontype: Number,
  promotionamount: Number,
  employee: Number,
  name: String,
	desc: String,
  code: String,
  date: String,
}

const Promotion: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [promotion, setPromotion] = React.useState<Partial<promotion>>({});

  const [promotiontype, setPromotiontype] = React.useState<EntPromotiontype[]>([]);
  const [promotionamount, setPromotionamount] = React.useState<EntPromotionamount[]>([]);
  const [employee, setEmployee] = React.useState<EntEmployee[]>([]);
 
  const getPromotiontype = async () => {
    const res = await http.listPromotiontype({ limit: 10, offset: 0 });
    setPromotiontype(res);
  };

  const getPromotionamount = async () => {
    const res = await http.listPromotionamount({ limit: 10, offset: 0 });
    setPromotionamount(res);
  };

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployee(res);
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
    getPromotiontype();
    getPromotionamount();
    getEmployee();
  }, []);

  // set data to object promotion
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Promotion;
    const { value } = event.target;
    setPromotion({ ...promotion, [name]: value });
    console.log(Promotion);
  };

  // clear input form
  function clear() {
    setPromotion({});
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
      case 'DESC':
        alertMessage("error","เงื่อนไขต้องกรอกไม่เกิน 30 ตัวอักษร");
        return;
      case 'CODE':
        alertMessage("error","รหัสต้องขึ้นต้น A-Z 2 ตัวและตามด้วย ตัวเลข 3 หลัก");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  // function save data
  function save() {
    promotion.date += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/promotions';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(promotion),
    };
    
    console.log(promotion); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  }

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
                  label="เพิ่มชื่อโปรโมชั่น"
                  name="name"
                  type="string"
                  value={promotion.name} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เงื่อนไขโปรโมชั่น</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เพิ่มเงื่อนไขโปรโมชั่น"
                  name="desc"
                  type="string"
                  value={promotion.desc || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทปรโมชั่น</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภทปรโมชั่น</InputLabel>
                <Select
                  name="promotiontype"
                  value={promotion.promotiontype || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {promotiontype.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.tYPE}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>จำนวนการใช้งานโปรโมชั่น</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกจำนวน</InputLabel>
                <Select
                  name="promotionamount"
                  value={promotion.promotionamount || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {promotionamount.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.aMOUNT}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่และเวลาหมดอายุ</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="exp."
                  name="date"
                  type="datetime-local"
                  value={promotion.date || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รหัสโปรโมชั่น</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="กรอกรหัสโปรโมชั่น"
                  name="code"
                  type="string"
                  value={promotion.code || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ผู้รับผิดชอบ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกผู้รับผิดชอบ</InputLabel>
                <Select
                  name="employee"
                  value={promotion.employee || ''} // (undefined || '') = ''
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


            <Grid item xs={5}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกโปรโมชั่น
              </Button>
            </Grid>

            <Grid item xs={4}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SearchIcon />}
                href="./Promotionsearch"
              >
                ค้นหาโปรโมชั่น
              </Button>
            </Grid>

            

          </Grid>

        </Container>

      </Content>

    </Page>
  );
};

export default Promotion;
