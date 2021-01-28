import React, { useState, useEffect } from 'react';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
    InfoCard,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';

import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { Alert } from '@material-ui/lab';

import { makeStyles, Theme, createStyles, ThemeProvider } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';

import SearchIcon from '@material-ui/icons/Search';

import { EntEmployee } from '../../api/models/EntEmployee';


// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(2),
   },
   insideLabel: {
    margin: theme.spacing(1),
  },
   button: {
    marginLeft: '40px',
  },
   textField: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
   },
    select: {
      width: 400 ,
      marginLeft:7,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    center: {
      marginTop: theme.spacing(2),
      marginLeft: theme.spacing(23),
    },
    cardtable: {
      marginTop: theme.spacing(2),
    },
    fieldText: {
      width: 200,
      marginLeft:7,
    },
    fieldLabel: {
      marginLeft:8,
      marginRight: 20,
    },
    table: {
        minWidth: 650,
    }
  }),
);

const searchcheck = {
    employeesearchcheck: true,
   
}

export default function SearchRoom() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [employee, setEmployee] = useState<EntEmployee[]>([]);
    

    const [employeesearch, setEmployeesearch] = useState(String);
    

  useEffect(() => {
    const getEmployees = async () => {
    const res = await api.listEmployee();
      setLoading(false);
      setEmployee(res);
    };
    getEmployees();
    }, [loading]);

    const SearchEmployee = async () => {
        const res = await api.listEmployee();
        const employeesearch = EmployeeSearch(res);
        
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setEmployee([]);
        if(employeesearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setEmployee(employeesearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.employeesearchcheck = true;
        
    }
    const EmployeeSearch = (res: any) => {
        const data = res.filter((filter: EntEmployee) => filter.eMPLOYEEID?.includes(employeesearch))
       // console.log(data.length)
        if (data.length != 0 && employeesearch != "") {
            return data;
        }
        else {
            searchcheck.employeesearchcheck = false;
            if(employeesearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const EmployeeSearchhandleChange = (event: any) => {
        setEmployeesearch(event.target.value as string);
    };
    

    return (
        <Page theme={pageTheme.service}>
            <Content>
            <InfoCard title="ค้นหาข้อมูลพนักงาน">

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>รหัสพนักงาน(EmployeeID)</strong></div>
                    <TextField
                        id="employeeid"
                       // label="ค้นหารหัสพนักงาน"
                        type="string"
                        size="medium"
                        value={employeesearch}
                        onChange={EmployeeSearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>

           

                    <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchEmployee();
                }}
                variant="contained"
                color="primary"
                startIcon={<SearchIcon />}
              >
                ค้นหา
             </Button>
             
        
             <div><br></br></div>
             <TableContainer component={Paper}>
                            <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                <TableCell align="center">รหัสพนักงาน</TableCell>
                                <TableCell align="center">ชื่อของพนักงาน</TableCell>
                                <TableCell align="center">ที่อยู่ของพนักงาน</TableCell>
                                <TableCell align="center">เลขบัตรประชาชน</TableCell>
                                <TableCell align="center">อายุ</TableCell>
                                <TableCell align="center">ตำแหน่ง</TableCell>
                                <TableCell align="center">เงินเดือน</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {employee.map((item:EntEmployee ) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.eMPLOYEEID}</TableCell>
                                    <TableCell align="center">{item.eMPLOYEENAME}</TableCell>
                                    <TableCell align="center">{item.eMPLOYEEADDRESS}</TableCell>
                                    <TableCell align="center">{item.iDCARDNUMBER}</TableCell>
                                    <TableCell align="center">{item.edges?.age?.aGE}</TableCell>
                                    <TableCell align="center">{item.edges?.position?.pOSITION}</TableCell>
                                    <TableCell align="center">{item.edges?.salary?.sALARY}</TableCell> 
                                </TableRow>
                                ))}
                            </TableBody>
                            </Table>
                        </TableContainer>

                        <div>
             {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} style={{ width: 400 ,marginTop: 20, marginLeft:6 }} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}</div>

                </InfoCard>

                        

            </Content>
     </Page>
    );
}