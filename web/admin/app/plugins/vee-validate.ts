/* eslint-disable camelcase */
import Vue from 'vue'
import { ValidationProvider, ValidationObserver, localize, extend } from 'vee-validate'
import ja from 'vee-validate/dist/locale/ja.json'
import { required, confirmed, max, min, email, image, alpha_dash, size } from 'vee-validate/dist/rules'
import { ValidationRule } from 'vee-validate/dist/types/types'

Vue.component('ValidationProvider', ValidationProvider)
Vue.component('ValidationObserver', ValidationObserver)

const hiraganaCustomRule: ValidationRule = {
  params: ['hiragana'],
  validate: (value: string): boolean => {
    const pattern: RegExp = /^[ぁ-ゔー]*$/g
    return pattern.test(value)
  },
  message: (field: string): string => {
    return `${field}はひらがなのみ使用できます`
  },
}

const passwordCustomRule: ValidationRule = {
  params: ['password'],
  validate: (value: string): boolean => {
    const pattern: RegExp = /^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$/g
    return pattern.test(value)
  },
  message: (field: string): string => {
    return `${field}は英数字と_!@#$_%^&*.?()-=+のみ使用できます`
  },
}

// Basic Rules
extend('required', { ...required })
extend('confirmed', { ...confirmed })
extend('max', { ...max })
extend('min', { ...min })
extend('email', { ...email })
extend('image', { ...image })
extend('alpha_dash', { ...alpha_dash })
extend('size', { ...size })

// Custom Rules
extend('hiragana', { ...hiraganaCustomRule })
extend('password', { ...passwordCustomRule })

localize('ja', ja)
