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
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import {
  EntAge,
  EntPosition,
  EntSalary,
} from '../../api/models';

// header css
const HeaderCustom = {
  minHeight: '50px',
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

interface employee {
  age: Number,
  position: Number,
  salary: Number,
  employeeid: String,
  employeename: String,
  employeeaddress: String,
  idcardnumber: String,
}

const employee: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [employee, setEmployee] = React.useState< Partial<employee>>({});

  const [age, setAge] = React.useState<EntAge[]>([]);
  const [position, setPosition] = React.useState<EntPosition[]>([]);
  const [salary, setSalary] = React.useState<EntSalary[]>([]);

  const getAge = async () => {
    const res = await http.listAge({ limit: 10, offset: 0 });
    setAge(res);
  };

  const getPosition = async () => {
    const res = await http.listPosition({ limit: 11, offset: 0 });
    setPosition(res);
  };

  const getSalary = async () => {
    const res = await http.listSalary({ limit: 21, offset: 0 });
    setSalary(res);
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
    getAge();
    getPosition();
    getSalary();
  }, []);

  // set data to object employee
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof employee;
    const { value } = event.target;
    setEmployee({ ...employee, [name]: value });
    console.log(employee);
  };

  // clear input form
  function clear() {
    setEmployee({});
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'EMPLOYEEID':
        alertMessage("error","รหัสพนักงานขึ้นต้นด้วย EM ตามด้วยเลข 2 ตัว");
        return;
      case 'EMPLOYEENAME':
        alertMessage("error","กรุณากรอกชื่อ");
        return;
      case 'EMPLOYEEADDRESS':
        alertMessage("error","กรุณากรอกที่อยู่");
        return;
      case 'IDCARDNUMBER':
        alertMessage("error","เลขประจำตัวประชาชน 13 หลัก");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }
  // function save data
  function save() {
    
    const apiUrl = 'http://localhost:8080/api/v1/employees';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(employee),
    };

    console.log(employee); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  };

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบเพิ่มข้อมูลพนักงาน`}>
        <div style={{ marginLeft: 10 }}>Team G16</div>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>รหัสของพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.root} noValidate autoComplete="off">
                <TextField label="รหัสพนักงาน" 
                  name = "employeeid"
                  type="string"
                  defaultValue=" " 
                  value={employee.employeeid ||''} // (undefined  '') = ''
                  onChange={handleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อของพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.root} noValidate autoComplete="off">
                <TextField label="ชื่อพนักงาน" 
                  name = "employeename"
                  type="string"
                  defaultValue=" " 
                  value={employee.employeename ||''} // (undefined  '') = ''
                  onChange={handleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ที่อยู่ของพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.root} noValidate autoComplete="off">
                <TextField label="ที่อยู่พนักงาน" 
                  name = "employeeaddress"
                  type="string"
                  defaultValue=" " 
                  value={employee.employeeaddress ||''} // (undefined  '') = ''
                  onChange={handleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เลขบัตรประชาชนของพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.root} noValidate autoComplete="off">
                <TextField label="เลขบัตรประชาชน" 
                  name = "idcardnumber"
                  type="string"
                  defaultValue=" " 
                  value={employee.idcardnumber ||''} // (undefined  '') = ''
                  onChange={handleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>อายุของพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกอายุของพนักงาน</InputLabel>
                <Select
                  name="age"
                  value={employee.age || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {age.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.aGE}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ตำแหน่งของพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกจำแหน่งของพนักงาน</InputLabel>
                <Select
                  name="position"
                  value={employee.position || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {position.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.pOSITION}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
            <div className={classes.paper}>เงินเดือนของพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกเงินเดือน</InputLabel>
                <Select
                  name="salary"
                  value={employee.salary || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {salary.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.sALARY}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
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
                บันทึกข้อมูล
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default employee;
