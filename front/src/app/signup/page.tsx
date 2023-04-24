'use client'
import { useState } from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { useForm } from 'react-hook-form';

import { signup, UserInfo } from 'features/user/signup';


const schema = yup.object().shape({
    name: yup.string().required('名前を入力してください'),
    email: yup.string().required('メールアドレスを入力してください').test('email_valdate', 'メールアドレスを確認してください',value => {
        if (value.includes('gmail.com')) return true;
        return false;
    }),
    password: yup.string().required('パスワードを入力してください')
})

export default function Signup({ item }: {item: string}) {
    const [ passState, setPassState ] = useState<boolean>(false)
    const { register, handleSubmit, formState: { errors } } = useForm<UserInfo>({
        resolver: yupResolver(schema)
    })

    const onSubmit = (data: UserInfo) => {
        console.log('処理実行')
        signup(data)
    }
    return (
        <div>
            <form
                onSubmit={handleSubmit(onSubmit)}
            >
                <input
                    type="text"
                    placeholder="ユーザ名を入力"
                    {...register('name')}
                />
                <input
                    type="text"
                    placeholder="メールアドレスを入力"
                    {...register('email')}
                />
                <input
                    type={passState ? 'text' : 'password'}
                    placeholder="パスワードを入力"
                    {...register('password')}
                />
                <input
                    type="button"
                    value="表示"
                    onClick={() => setPassState(!passState)}
                />
                <button>
                    送信
                </button>
                <p>{ item }</p>
                <>
                    { errors.name &&
                        errors.name.message
                    }
                    {errors.email &&
                        errors.email.message
                    }
                    {errors.password &&
                        errors.password.message
                    }
                </>
            </form>
        </div>
    )
}