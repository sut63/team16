import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import SearchIcon from '@material-ui/icons/Search'; //search icon
import PowerSettingsNewIcon from '@material-ui/icons/PowerSettingsNew'; //log off icon
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
  EntCourse,
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

// alert setting
const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
});

interface bookcourse {
  member: Number,
  employee: Number,
  course: Number,
  booktime: String,
  phone : String,
  access : Number,
  detail : String,
}

const Bookcourse: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [bookcourse, setBookcourse] = React.useState<Partial<bookcourse>>({});
  const [course, setCourse] = React.useState<EntCourse[]>([]);
  const [member, setMember] = React.useState<EntMember[]>([]);
  const [employee, setEmployee] = React.useState<EntEmployee[]>([]);
  const [phoneError, setphoneError] = React.useState('');
  const [detailError, setdetailError] = React.useState('');
  const [accessError, setaccessError] = React.useState('');

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployee(res);
  };

  const getCourse = async () => {
    const res = await http.listCourse({ limit: 10, offset: 0 });
    setCourse(res);
  };

  const getMember = async () => {
    const res = await http.listMember({ limit: 10, offset: 0 });
    setMember(res);
  };
  
  // Lifecycle Hooks
  useEffect(() => {
    getCourse();
    getEmployee();
    getMember();
  }, []);

  // set data to object bookcourse
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof Bookcourse;
    const { value } = event.target;

    const validateValue = value.toString()
    checkPattern(name, validateValue)
    
    setBookcourse({ ...bookcourse, [name]: value });
    console.log(bookcourse);
  };

// ฟังก์ชั่นสำหรับ validate รายละเอียดการจอง
const validateldetail = (val: string) => {
  return val.charCodeAt(0) >= 65 && val.charCodeAt(0) <= 90 ? true : false;
}

// ฟังก์ชั่นสำหรับ validate จำนวณผู้เข้าจอง
const validatelaccess = (val: string) => {
  return val.length == 1 ? true : false;
}

// ฟังก์ชั่นสำหรับ validate เลขโทรศัพท์
const validatetelphone = (val: string) => {
  return val.length == 10 ? true : false;
}


// สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
const checkPattern  = (id: string, value: string) => {
  switch(id) {
    case 'detail':
      validateldetail(value) ? setdetailError('') : setdetailError('รายละเอียดการจอง30ตัวอักษร');
      return;
    case 'access':
      validatelaccess(value) ? setaccessError('') : setaccessError('จำนวณผู้เข้าใช้สนามไม่ถูกต้อง');
      return;
    case 'phone':
      validatetelphone(value) ? setphoneError('') : setphoneError('หมายเลขโทรศัพท์ตัวเลข 10 หลัก')
      return;
      
    default:
      return;
  }
}

const alertMessage = (icon: any, title: any) => {
  Toast.fire({
    icon: icon,
    title: title,
  });
}

const checkCaseSaveError = (field: string) => {
  switch(field) {
    case 'DETAIL':
      alertMessage("error","รายละเอียดการจอง30ตัวอักษร");
      return;
    case 'ACCESS':
      alertMessage("error","จำนวณผู้เข้าใช้ไม่ถูกต้อง");
      return;
    case 'PHONE':
      alertMessage("error","หมายเลขโทรศัพท์ตัวเลข 10 หลัก");
      return;
    default:
      alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
      return;
  }
}

  // clear input form
  function clear() {
    setBookcourse({});
  }

  // function save data
  function save() {
    bookcourse.booktime += ":00+07:00";
    bookcourse.access = Number(bookcourse.access);
    const apiUrl = 'http://localhost:8080/api/v1/bookcourses';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(bookcourse),
    };

    console.log(bookcourse); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
      <Header style={HeaderCustom} title={`ระบบจองสนามกีฬา`}>
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
                  value={bookcourse.member || ''} // (undefined || '') = ''
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
              <div className={classes.paper}>สนามที่ต้องการจอง</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกสนาม</InputLabel>
                <Select
                  name="course"
                  value={bookcourse.course || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {course.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.cOURSE}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เบอร์ติดต่อ</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เพิ่มเบอร์ติดต่อ"
                  name="phone"
                  type="string"
                  value={bookcourse.phone || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>จำนวณผู้เข้าใช้</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เพิ่มจำนวณผู้เข้าใช้"
                  name="access"
                  type="number"
                  value={bookcourse.access || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อพนักงานที่บันทึก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพนักงาน</InputLabel>
                <Select
                  name="employee"
                  value={bookcourse.employee || ''} // (undefined || '') = ''
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
              <div className={classes.paper}>รายละเอียดการเข้าใช้</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เพิ่มรายละเอียดการเข้าใช้"
                  name="detail"
                  type="string"
                  value={bookcourse.detail || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ต้องการจอง</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันที่"
                  name="booktime"
                  type="datetime-local"
                  value={bookcourse.booktime || ''} // (undefined || '') = ''
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
                บันทึกการจองสนาม
              </Button>
            </Grid>
            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SearchIcon />}
                href="./Bookcoursesearch"
              >
                ค้นหาการจองสนาม
              </Button>
            </Grid>


          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Bookcourse;
