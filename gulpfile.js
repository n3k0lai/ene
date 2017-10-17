const gulp = require('gulp')
const babel = require('gulp-babel')
const watch = require('gulp-watch')

const src = 'src/*.js'
const dest = 'lib/'

gulp.task('build', () => {
  return gulp.src(src)
    .pipe(babel())
    .pipe(gulp.dest(dest))
})

gulp.task('watch', [ 'build' ], () => {
  watch(src, () => {
    gulp.start('build')
  })
})

gulp.task('default', [ 'build' ])
