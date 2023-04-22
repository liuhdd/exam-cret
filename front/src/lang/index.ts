import { createI18n } from 'vue-i18n'
import  zhCnLocale  from '@/lang/package/zh-cn'


const messages = {
    'zh-cn': {
      ...zhCnLocale
    }
};


const i18n = createI18n({
    legacy: false,
    locale: 'zh-cn',
    messages: messages,
    globalInjection: true
})

export default i18n
