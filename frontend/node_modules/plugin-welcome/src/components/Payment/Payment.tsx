import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
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
import { DefaultApi } from '../../api/apis/'; // Api Gennerate From Command
import {
  EntMember,
  EntEmployee,
  EntPaymenttype,
  EntPromotion,
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

interface payment {
  member: Number,
  employee: Number,
  promotion: Number,
  paymenttype: Number,
  paymentamount: String,
  paymentdate: String,
}

const Payment: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [payment, setPayment] = React.useState<Partial<payment>>({});

  const [member, setMember] = React.useState<EntMember[]>([]);
  const [promotion, setPromotion] = React.useState<EntPromotion[]>([]);
  const [paymenttype, setPaymenttype] = React.useState<EntPaymenttype[]>([]);
  const [employee, setEmployee] = React.useState<EntEmployee[]>([]);

 
  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployee(res);
  };

  const getMember = async () => {
    const res = await http.listMember({ limit: 10, offset: 0 });
    setMember(res);
  };

  const getPromotion = async () => {
    const res = await http.listPromotion({ limit: 10, offset: 0 });
    setPromotion(res);
  };

  const getPaymenttype = async () => {
    const res = await http.listPaymenttype({ limit: 10, offset: 0 });
    setPaymenttype(res);
  };
  
  // Lifecycle Hooks
  useEffect(() => {
    getEmployee();
    getMember();
    getPaymenttype();
    getPromotion();
  }, []);

  // set data to object payment
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Payment;
    const { value } = event.target;
    setPayment({ ...payment, [name]: value });
    console.log(payment);
  };

  // clear input form
  function clear() {
    setPayment({});
  }

  // function save data
  function save() {
    payment.paymentdate += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/payments';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payment),
    };

    // alert setting
    const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
    });
    
    console.log(payment); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => {
        console.log(response.json());
        if (response.ok === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบชำระเงิน`}>
      </Header>

      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อ Member</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือก Member Name</InputLabel>
                <Select
                  name="member"
                  value={payment.member || ''} // (undefined || '') = ''
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

            <Grid item xs={3}>
              <div className={classes.paper}>โปรโมชั่น</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกโปรโมชั่น</InputLabel>
                <Select
                  name="promotion"
                  value={payment.promotion || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {promotion.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nAME}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ช่องทางการชำระ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภทการชำระ</InputLabel>
                <Select
                  name="paymenttype"
                  value={payment.paymenttype || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {paymenttype.map(item => {
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
              <div className={classes.paper}>ชื่อพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพนักงาน</InputLabel>
                <Select
                  name="employee"
                  value={payment.employee || ''} // (undefined || '') = ''
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

            <Grid item xs={3}>
              <div className={classes.paper}>ยอดที่ชำระเงิน</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.root} noValidate autoComplete="off">
                <TextField label="ราคา" 
                  name = "paymentamount"
                  type="string"
                  defaultValue=" " 
                  value={payment.paymentamount || ''} // (undefined || '') = ''
                  onChange={handleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ชำระเงิน</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันที่"
                  name="paymentdate"
                  type="datetime-local"
                  value={payment.paymentdate || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกการชำระเงิน
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Payment;
