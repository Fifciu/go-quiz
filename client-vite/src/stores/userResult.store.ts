import { defineStore } from 'pinia';
import { Answer, client, QuestionWithAnswers, UserResult, FinalUserResult } from 'api-client';

interface QuestionWithAnswersAndIndicator extends QuestionWithAnswers {
  isAnswered?: boolean
};
interface UserResultStoreState extends UserResult {
  questions: QuestionWithAnswersAndIndicator[],
  answers: Answer[],
  results: FinalUserResult[]
};

export const useUserResultStore = defineStore('userResult', {
  state: (): UserResultStoreState => ({
    id: -1, 
    test_id: -1,
    user_id: -1,
    start_datetime: '',
    finish_datetime: '',
    questions: [],
    answers: [],
    results: []
  }),
  getters: {
    isAlreadyFinished (state) {
      return state.answers.length === state.questions.length
    },
    getNextQuestion (state) {
      if (!this.isAlreadyFinished) {
        return state.questions[state.answers.length]
      }
      return -1;
    }
  },
  actions: {
    async start(testId: number) {
      if (!testId && this.id > 0) {
        return;
      }
      const results = await client.results.start(testId);
      this.id = results.id;
      this.test_id = results.test_id;
      this.user_id = results.user_id;
      this.start_datetime = results.start_datetime;
    },

    async getQuestions() {
      if (!this.test_id) {
        throw new Error(`[useUserStore] App didn't call "start" method!`)
      }
      this.questions = await client.tests.questionsAndAnswers(this.test_id);
    },

    async answer(answerId: number) {
      if (!this.test_id) {
        throw new Error(`[useUserStore] App didn't call "start" method!`)
      }
      const questionContainingAnswerIndex = this.questions.findIndex(question => question.answers.some(answer => answer.id == answerId));
      const questionContainingAnswer = this.questions[questionContainingAnswerIndex];
      if (!questionContainingAnswer) {
        throw new Error(`[useUserStore] Answer not matching any question!`)
      }
      if (questionContainingAnswer.isAnswered) {
        throw new Error(`[useUserStore] Question already answered`)
      }

      await client.answers.do(answerId);
      this.answers.push(
        questionContainingAnswer.answers.find(answer => answer.id = answerId) as Answer
      );
      questionContainingAnswer.isAnswered = true;
      return 
    },

    async finish () {
      if (!this.test_id) {
        throw new Error(`[useUserStore] App didn't call "start" method!`)
      }
      if (!this.isAlreadyFinished) {
        throw new Error(`[useUserStore] Cannot finish. User didn't answer every question!`)
      }
      this.results = await client.results.finish(this.id);
    },

    clearStorage () {
      this.id = -1;
      this.test_id = -1;
      this.user_id = -1;
      this.start_datetime = '';
      this.finish_datetime = '';
      this.questions = [];
      this.answers = [];
      this.results = []
    }
  }
});
