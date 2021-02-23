import React, { useEffect, useState } from 'react';
import { createHttpClient } from '../../http/client';
import { Loader } from '../Loader';
import { Alert } from '../Alert';
import { EmailForm } from './EmailForm';
import { MPlanForm } from './MPlanForm';
import {
  NotificationTimeSpan,
  NotificationTimeSpanForm,
} from './NotificationTimeSpanForm';

type AlertState = {
  visible: boolean;
  kind: string;
  message: string;
};

type NotificationTimeSpanState = {
  timeSpans: NotificationTimeSpan[];
  editable: boolean;
};

export const SettingPage: React.FC<{}> = () => {
  const [loading, setLoading] = useState(false);
  const [email, setEmail] = useState('');
  const [alert, setAlert] = useState<AlertState>({
    visible: false,
    kind: '',
    message: '',
  });
  const [
    notificationTimeSpanState,
    setNotificationTimeSpanState,
  ] = useState<NotificationTimeSpanState>({
    timeSpans: [],
    editable: false,
  });

  useEffect(() => {
    setLoading(true);
    const client = createHttpClient();
    client
      .get('/api/v1/me')
      .then((response) => {
        console.log(response.data);
        const timeSpans = response.data['notificationTimeSpans']
          ? response.data['notificationTimeSpans']
          : [];
        setEmail(response.data['email']);
        setNotificationTimeSpanState({
          ...notificationTimeSpanState,
          timeSpans: timeSpans,
        });
      })
      .catch((error) => {
        console.log(error);
        handleShowAlert('danger', 'システムエラーが発生しました');
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  const handleShowAlert = (kind: string, message: string) => {
    setAlert({ visible: true, kind: kind, message: message });
  };

  const handleHideAlert = () => {
    setAlert({ ...alert, visible: false });
  };

  const handleOnChangeEmail = (
    event: React.ChangeEvent<HTMLInputElement>
  ): void => {
    setEmail(event.currentTarget.value);
  };

  const handleUpdateEmail = (email: string): void => {
    const client = createHttpClient();
    client
      .post('/api/v1/me/email', {
        email: email,
      })
      .then((response) => {
        handleShowAlert('success', 'メールアドレスを更新しました！');
      })
      .catch((error) => {
        console.log(error);
        if (error.response.status === 400) {
          handleShowAlert('danger', '正しいメールアドレスを入力してください');
        } else {
          // TODO: external message
          handleShowAlert('danger', 'システムエラーが発生しました');
        }
      });
  };

  const handleSetTimeSpanEditable = (value: boolean) => {
    setNotificationTimeSpanState({
      ...notificationTimeSpanState,
      editable: value,
    });
  };

  const handleAddTimeSpan = () => {
    if (notificationTimeSpanState.timeSpans.length === 3) {
      return;
    }
    setNotificationTimeSpanState({
      ...notificationTimeSpanState,
      timeSpans: [
        ...notificationTimeSpanState.timeSpans,
        { fromHour: 0, fromMinute: 0, toHour: 0, toMinute: 0 },
      ],
    });
  };

  const handleDeleteTimeSpan = (index: number) => {
    let timeSpans = notificationTimeSpanState.timeSpans.slice();
    if (index >= timeSpans.length) {
      return;
    }
    timeSpans.splice(index, 1);
    setNotificationTimeSpanState({
      ...notificationTimeSpanState,
      timeSpans: timeSpans,
    });
  };

  const handleOnChangeTimeSpan = (
    name: string,
    index: number,
    value: number
  ) => {
    let timeSpans = notificationTimeSpanState.timeSpans.slice();
    timeSpans[index][name as keyof NotificationTimeSpan] = value;
    setNotificationTimeSpanState({
      ...notificationTimeSpanState,
      timeSpans: timeSpans,
    });
  };

  const handleUpdateTimeSpan = () => {
    const timeSpans: NotificationTimeSpan[] = [];
    for (const timeSpan of notificationTimeSpanState.timeSpans) {
      for (const [k, v] of Object.entries(timeSpan)) {
        timeSpan[k as keyof NotificationTimeSpan] = v;
      }
      if (
        timeSpan.fromHour === 0 &&
        timeSpan.fromMinute === 0 &&
        timeSpan.toHour === 0 &&
        timeSpan.toMinute === 0
      ) {
        // Ignore zero value
        continue;
      }
      timeSpans.push(timeSpan);
    }

    const client = createHttpClient();
    client
      .post('/api/v1/me/notificationTimeSpan', {
        notificationTimeSpans: timeSpans,
      })
      .then((response) => {
        handleShowAlert('success', 'レッスン希望時間帯を更新しました！');
      })
      .catch((error) => {
        console.log(error);
        if (error.response.status === 400) {
          handleShowAlert(
            'danger',
            '正しいレッスン希望時間帯を選択してください'
          );
        } else {
          // TODO: external message
          handleShowAlert('danger', 'システムエラーが発生しました');
        }
      });

    setNotificationTimeSpanState({
      editable: false,
      timeSpans: timeSpans,
    });
  };

  return (
    <div>
      <h1 className="page-title">設定</h1>
      {loading ? (
        <Loader
          loading={loading}
          message={'Loading data ...'}
          css={'background: rgba(255, 255, 255, 0)'}
          size={50}
        />
      ) : (
        <>
          <Alert handleCloseAlert={handleHideAlert} {...alert} />
          <EmailForm
            email={email}
            handleOnChange={handleOnChangeEmail}
            handleUpdateEmail={handleUpdateEmail} // TODO: inline function
          />
          <NotificationTimeSpanForm
            handleAdd={handleAddTimeSpan}
            handleDelete={handleDeleteTimeSpan}
            handleUpdate={handleUpdateTimeSpan}
            handleOnChange={handleOnChangeTimeSpan}
            handleSetEditable={handleSetTimeSpanEditable}
            {...notificationTimeSpanState}
          />
          {/*<MPlanForm {...this.state.mPlan} />*/}
        </>
      )}
    </div>
  );
};