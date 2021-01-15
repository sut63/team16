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
  EntClassifier,
  EntEmployee,
  EntZone,
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

interface equipment {
  EQUIPMENTCLASSIFIER: Number,
  EMPLOYEE: Number,
  EQUIPMENTTYPE: Number,
  EQUIPMENTZONE: Number,
  EQUIPMENTNAME: String,
  EQUIPMENTAMOUNT: Number,
  EQUIPMENTDATE: String,
  EQUIPMENTDETAIL: String,
}

const Equipment: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [equipment, setEquipment] = React.useState<Partial<equipment>>({});
  const [classifier, setClassifier] = React.useState<EntClassifier[]>([]);
  const [zone, setZone] = React.useState<EntZone[]>([]);
  const [equipmenttype, setEquipmenttype] = React.useState<EntEquipmenttype[]>([]);
  const [employee, setEmployee] = React.useState<EntEmployee[]>([]);

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployee(res);
  };

  const getClassifier = async () => {
    const res = await http.listClassifier({ limit: 10, offset: 0 });
    setClassifier(res);
  };

  const getEquipmenttype = async () => {
    const res = await http.listEquipmenttype({ limit: 10, offset: 0 });
    setEquipmenttype(res);
  };

  const getZone = async () => {
    const res = await http.listZone({ limit: 10, offset: 0 });
    setZone(res);
  };
  
  // Lifecycle Hooks
  useEffect(() => {
    getClassifier();
    getEmployee();
    getEquipmenttype();
    getZone();
  }, []);

  // set data to object equipment
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Equipment;
    const { value } = event.target;
    setEquipment({ ...equipment, [name]: value });
    console.log(equipment);
  };

  // clear input form
  function clear() {
    setEquipment({});
  }
  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
  });
  

  // function save data
  function save() {
    if (equipment.EQUIPMENTAMOUNT <= 0){
      Toast.fire({
        icon: 'error',
        title: 'บันทึกข้อมูลไม่สำเร็จ',
      });
      return 0;
    }

    equipment.EQUIPMENTDATE += ":00+07:00";
    equipment.EQUIPMENTAMOUNT = parseInt(equipment.EQUIPMENTAMOUNT);
    const apiUrl = 'http://localhost:8080/api/v1/equipments';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(equipment),
    };

    
    console.log(equipment); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
      <Header style={HeaderCustom} title={`ระบบเพิ่มอุปกรณ์`}>
      </Header>

      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อของอุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.root} noValidate autoComplete="off">
                <TextField label="=ชื่ออุปกรณ์" 
                  name = "EQUIPMENTNAME"
                  type="string"
                  defaultValue=" " 
                  value={equipment.EQUIPMENTNAME || ''} // (undefined || '') = ''
                  onChange={handleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทของอุปกรณ์</div>
            </Grid>

            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภทของอุปกรณ์</InputLabel>
                <Select
                  name="EQUIPMENTTYPE"
                  value={equipment.EQUIPMENTTYPE || ''} // (undefined || '') = ''
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
              <div className={classes.paper}>จำนวนของอุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.root} noValidate autoComplete="off">
                <TextField label="=จำนวนอุปกรณ์" 
                  name = "EQUIPMENTAMOUNT"
                  type="number"
                  defaultValue=" " 
                  value={equipment.EQUIPMENTAMOUNT || ''} // (undefined || '') = ''
                  onChange={handleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>ลักษณะนามของอุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกลักษณะนาม</InputLabel>
                <Select
                  name="EQUIPMENTCLASSIFIER"
                  value={equipment.EQUIPMENTCLASSIFIER || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {classifier.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.eQUIPMENTCLASSIFIER}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ตำแหน่งการเก็บอุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกตำแหน่ง</InputLabel>
                <Select
                  name="EQUIPMENTZONE"
                  value={equipment.EQUIPMENTZONE || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {zone.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.eQUIPMENTZONE}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รายละเอียดของอุปกรณ์</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.root} noValidate autoComplete="off">
                <TextField label="=รายละเอียด" 
                  name = "EQUIPMENTDETAIL"
                  type="string"
                  defaultValue=" " 
                  value={equipment.EQUIPMENTDETAIL || ''} // (undefined || '') = ''
                  onChange={handleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพนักงาน</InputLabel>
                <Select
                  name="EMPLOYEE"
                  value={equipment.EMPLOYEE || ''} // (undefined || '') = ''
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
              <div className={classes.paper}>วันที่เพิ่ม</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันที่"
                  name="EQUIPMENTDATE"
                  type="datetime-local"
                  value={equipment.EQUIPMENTDATE || ''} // (undefined || '') = ''
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
                บันทึกการเพิ่มอุปกรณ์
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Equipment;
