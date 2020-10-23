import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import InputLabel from '@material-ui/core/InputLabel';
import { EntScholarshiptype } from '../../api/models/EntScholarshiptype';
import { EntEducationlevel } from '../../api/models/EntEducationlevel';
import { EntStudyplan } from '../../api/models/EntStudyplan';
import { EntSemester } from '../../api/models/EntSemester';
import { EntUser } from '../../api';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'left',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ระบบรับข้อมูลจากผู้ให้ทุนการศึกษา' };
  const api = new DefaultApi();

  const [scholarshiptypes, setScholarshiptype] = useState<EntScholarshiptype[]>([]);
  const [studyplans, setStudyplan] = useState<EntStudyplan[]>([]);
  const [semesters, setSemester] = useState<EntSemester[]>([]);
  const [educationlevels, setEducationlevel] = useState<EntEducationlevel[]>([]);
  const [users, setUser] = useState<EntUser[]>([]);
  const [loading, setLoading] = useState(true);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [studyplanid, setStudyplanID] = useState(Number);
  const [educationlevelid, setEducationlevelID] = useState(Number);
  const [semesterid, setSemesterID] = useState(Number);
  const [userid, setUserID] = useState(Number);
  const [scholarshiptypeid, setscholarshiptypeName] = useState(Number);
  const [organization, setOrganization] = useState(String);

  const [scholarshipname, setScholarshipName] = useState(String);

  useEffect(() => {
    const getScholarshiptype = async () => {
      const s = await api.listScholarshiptype({ limit: 10, offset: 0 });
      setScholarshiptype(s);
    };
    getScholarshiptype();

    const getStudyplan = async () => {
      const p = await api.listStudyplan({ limit: 10, offset: 0 });
      setStudyplan(p);
    };
    getStudyplan();

    const getEducationlevel = async () => {
      const p = await api.listEducationlevel({ limit: 10, offset: 0 });
      setEducationlevel(p);
    };
    getEducationlevel();

    const getSemester = async () => {
      const st = await api.listSemester({ limit: 10, offset: 0 });
      setSemester(st);
    };
    getSemester();

    const getUser = async () => {
      const us = await api.listUser({ limit: 10, offset: 0 });
      setUser(us);
    };
    getUser();

  }, [loading]);

  const CreateScholarship = async () => {
    const scholarship = {
      educationLevelID: educationlevelid,
      organization: organization,
      scholarshipName: scholarshipname,
      scholarshipTypeID: scholarshiptypeid,
      semesterID: semesterid,
      studyPlanID: studyplanid,
      userID: userid,
    }

    const res: any = await api.createScholarship({ scholarship: scholarship });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  const scholarshiptypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setscholarshiptypeName(event.target.value as number);
  };

  const educationlevelhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEducationlevelID(event.target.value as number);
  };

  const studyplanhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStudyplanID(event.target.value as number);
  };

  const semesterNamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSemesterID(event.target.value as number);
  };

  const ScholarshipNamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setScholarshipName(event.target.value as string);
  };

  const OrganizationhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setOrganization(event.target.value as string);
  };
  const UserhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserID(event.target.value as number);
  };
  
  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`Welcome ${profile.givenName || 'ระบบรับข้อมูลจากผู้ให้ทุนการศึกษา'}`}
      //  subtitle = {`${username.givenuser}`}
      ></Header>
      <Content>
        <ContentHeader title="กรุณากรอกข้อมูล">
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

          <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="user">ผู้ให้ทุน</InputLabel>
                <Select
                  label="ผู้ให้ทุน"
                  id="user"
                  value={userid}
                  onChange={UserhandleChange}
                  style={{ width: 300 }}
                >

                  {users.map((item: EntUser) =>
                    <MenuItem value={item.id}>{item.username}</MenuItem>)}

                </Select>
              </FormControl>
            </div>


            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="ScholarshipName"></InputLabel>
                <TextField
                  id="ScholarshipName"
                  label=" ชื่อทุนการศึกษา"
                  type="string"
                  value={scholarshipname}
                  onChange={ScholarshipNamehandleChange}
                  className={classes.textField}
                // InputLabelProps={{
                //   shrink: true,
                // }}

                />
              </FormControl>
            </div>


            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="Organization"></InputLabel>
                <TextField
                  id="Organization"
                  label="ชื่อองค์กร/สถาบันการศึกษา"
                  type="string"
                  value={organization}
                  onChange={OrganizationhandleChange}
                  className={classes.textField}
                // InputLabelProps={{
                //   shrink: true,
                // }}

                />
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="scholarship">ประเภททุน</InputLabel>
                <Select
                  label="ประเภททุน"
                  id="scholarship"
                  value={scholarshiptypeid}
                  onChange={scholarshiptypehandleChange}
                  style={{ width: 300 }}
                >

                  {scholarshiptypes.map((item: EntScholarshiptype) =>
                    <MenuItem value={item.id}>{item.scholarshiptypename}</MenuItem>)}

                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="education">ระดับการศึกษา</InputLabel>
                <Select
                  label="ระดับการศึกษา"
                  id="education"
                  value={educationlevelid}
                  onChange={educationlevelhandleChange}
                  style={{ width: 300 }}
                >
                  {educationlevels.map((item: EntEducationlevel) =>
                    <MenuItem value={item.id}>{item.educationlevelname}</MenuItem>)}
                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="studyplan">แผนการเรียน/สาขาวิชา</InputLabel>
                <Select
                  label="แผนการเรียน/สาขาวิชา"
                  id="studyplan"
                  value={studyplanid}
                  onChange={studyplanhandleChange}
                  style={{ width: 300 }}
                >
                  {studyplans.map((item: EntStudyplan) =>
                    <MenuItem value={item.id}>{item.studyplanname}</MenuItem>)}
                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="semester">ภาคการศึกษา</InputLabel>
                <Select
                  label="ภาคการศึกษา"
                  id="semester"
                  value={semesterid}
                  onChange={semesterNamehandleChange}
                  style={{ width: 300 }}
                >
                  {semesters.map((item: EntSemester) =>
                    <MenuItem value={item.id}>{item.semestername}</MenuItem>)}
                </Select>
              </FormControl>
            </div>

           


            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateScholarship();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยัน
             </Button>
              <Button
                style={{
                  marginLeft: 20,
                  backgroundColor: 'red'
                }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                ยกเลิก
             </Button>
            </div>

          </form>
        </div>
      </Content>
    </Page>
  );
}