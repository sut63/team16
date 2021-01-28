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
  EntMember,
  EntEmployee,
  EntEquipment,
  EntEquipmenttype,
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

interface equipmentrental {
  member: Number,
  employee: Number,
  equipment: Number,
  equipmenttype: Number,
  rentalamount: Number,
  rentaldate: String,
  returndate: String,
}

const Equipmentrental: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [equipmentrental, setEquipmentrental] = React.useState<Partial<equipmentrental>>({});

  const [member, setMember] = React.useState<EntMember[]>([]);
  const [equipment, setEquipment] = React.useState<EntEquipment[]>([]);
  const [equipmenttype, setEquipmenttype] = React.useState<EntEquipmenttype[]>([]);
  const [employee, setEmployee] = React.useState<EntEmployee[]>([]);

 
  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployee(res);
  };

  const getMember = async () => {
    const res = await http.listMember({ limit: 10, offset: 0 });
    setMember(res);
  };

  const getEquipment = async () => {
    const res = await http.listEquipment({ limit: 10, offset: 0 });
    setEquipment(res);
  };

  const getEquipmenttype = async () => {
    const res = await http.listEquipmenttype({ limit: 10, offset: 0 });
    setEquipmenttype(res);
  };
  
  // Lifecycle Hooks
  useEffect(() => {
    getEmployee();
    getMember();
    getEquipmenttype();
    getEquipment();
  }, []);

  // set data to object payment
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Equipmentrental;
    const { value } = event.target;
    setEquipmentrental({ ...equipmentrental, [name]: value });
    console.log(equipmentrental);
  };

  // clear input form
  function clear() {
    setEquipmentrental({});
  }

 // alert setting
 const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
});

const alertMessage = (icon: any, title: any) => {
  Toast.fire({
    icon: icon,
    title: title,
  });
}

  // function save data
  function save() {
    equipmentrental.rentaldate += ":00+07:00";
    equipmentrental.returndate += ":00+07:00";
    equipmentrental.rentalamount = Number(equipmentrental.rentalamount);
    const apiUrl = 'http://localhost:8080/api/v1/equipmentrentals';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(equipmentrental),
    };

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
          alertMessage("error", data.error);
        }
      });
  };
   
    
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบยืมอุปกรณ์กีฬา`}>
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
              <div className={classes.paper}>ชื่อผู้ยืม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกชื่อผู้ยืม</InputLabel>
                <Select
                  name="member"
                  value={equipmentrental.member || ''} // (undefined || '') = ''
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
              <div className={classes.paper}>ชื่ออุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกชื่ออุปกรณ์</InputLabel>
                <Select
                  name="equipment"
                  value={equipmentrental.equipment || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {equipment.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.eQUIPMENTNAME}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ชนิดอุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภทอุปกรณ์</InputLabel>
                <Select
                  name="equipmenttype"
                  value={equipmentrental.equipmenttype || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {equipmenttype.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.eQUIPMENTTYPE}
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
                  value={equipmentrental.employee || ''} // (undefined || '') = ''
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
              <div className={classes.paper}>จำนวนอุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.root} noValidate autoComplete="off">
                <TextField label="จำนวน" 
                  name = "rentalamount"
                  type="Number"
                  defaultValue=" " 
                  value={equipmentrental.rentalamount || ''} // (undefined || '') = ''
                  onChange={handleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ยืมอุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันที่"
                  name="rentaldate"
                  type="datetime-local"
                  value={equipmentrental.rentaldate || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ครบกำหนดการยืม</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันที่"
                  name="returndate"
                  type="datetime-local"
                  value={equipmentrental.returndate || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={5}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกการยืมอุปกรณ์
              </Button>
            </Grid>

            <Grid item xs={4}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SearchIcon />}
                href="./Equipmentrentalsearch"
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

export default Equipmentrental;
