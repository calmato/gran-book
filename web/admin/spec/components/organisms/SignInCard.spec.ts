import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SignInCard from '~/components/organisms/SignInCard.vue'

describe('components/organisms/SignInCard', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(SignInCard, { ...Options })
  })

  it('test is not exist', () => {
    expect(wrapper.html()).toMatchSnapshot()
  })
})
