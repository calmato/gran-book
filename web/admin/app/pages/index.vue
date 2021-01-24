<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row>
        <v-col v-for="card in cards" :key="card" cols="12">
          <v-card>
            <v-subheader>{{ card }}</v-subheader>

            <v-list two-line>
              <template v-for="n in 6">
                <v-list-item :key="n">
                  <v-list-item-avatar color="grey darken-1"> </v-list-item-avatar>

                  <v-list-item-content>
                    <v-list-item-title>Message {{ n }}</v-list-item-title>

                    <v-list-item-subtitle>
                      Lorem ipsum dolor sit amet, consectetur adipisicing elit. Nihil repellendus distinctio similique
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>

                <v-divider v-if="n !== 6" :key="`divider-${n}`" inset></v-divider>
              </template>
            </v-list>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  setup(_props, { root }: SetupContext) {
    const router = root.$router
    const cards: string[] = ['Today', 'Yesterday']

    const handleClick = (): void => {
      router.push('/signin')
    }

    return {
      cards,
      handleClick,
    }
  },
})
</script>
