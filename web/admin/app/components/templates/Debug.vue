<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row>
        <v-col cols="12">
          <v-card>
            <v-card-title>UID</v-card-title>
            <v-card-text>{{ userId }}</v-card-text>
            <v-card-title>Token</v-card-title>
            <v-card-text class="token">
              {{ omitToken }}
              <v-tooltip v-model="showTooltip" top>
                <template v-slot:activator="{}">
                  <v-btn icon @click="handleCopyText">
                    <v-icon>mdi-clipboard-text</v-icon>
                  </v-btn>
                </template>
                <span>コピーしました</span>
              </v-tooltip>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from '@nuxtjs/composition-api'
import useClipboard from 'vue-clipboard3'

export default defineComponent({
  props: {
    userId: {
      type: String,
      reduired: false,
      default: '',
    },
    token: {
      type: String,
      required: false,
      default: '',
    },
  },

  setup(props) {
    const { toClipboard } = useClipboard()

    const showTooltip = ref<boolean>(false)

    const omitToken = computed(() => {
      const length: number = 64
      const omission: string = '...'

      if (props.token.length < length) {
        return props.token
      }

      return props.token.substring(0, length) + omission
    })

    const handleCopyText = async () => {
      await toClipboard(props.token)
        .then(() => {
          showTooltip.value = true
          setTimeout(hiddenTooltip, 2000)
        })
        .catch((err) => {
          console.error(err)
        })
    }

    const hiddenTooltip = () => {
      showTooltip.value = false
    }

    return {
      showTooltip,
      omitToken,
      handleCopyText,
    }
  },
})
</script>

<style scoped>
p.token {
  word-break: break-all;
}
</style>
