import {
  AfterViewInit,
  Component,
  ElementRef,
  EventEmitter,
  Input,
  OnDestroy,
  OnInit,
  Output,
  Renderer2,
  ViewChild
} from '@angular/core';
import { getOffsetValueFromLine, isOffsetLine } from '../../../shared/lib/tscript';

@Component({
  selector: 'app-editor',
  templateUrl: './editor.component.html',
  styleUrls: ['./editor.component.scss']
})
export class EditorComponent implements OnInit, OnDestroy, AfterViewInit {

  _textContent = '';

  @Input()
  set textContent(f: string) {
    this._textContent = f;
    if (this.editableContent) {
      this.updateInnerHtml(this._textContent);
    }
  }

  get textContent(): string {
    if (this.editableContent) {
      return this.editableContent.nativeElement.innerText;
    }
    return this._textContent;
  }

  @Input()
  readonly: boolean = false;

  @Input()
  wrap: boolean = false;

  @Output()
  textContentChange: EventEmitter<string> = new EventEmitter<string>();

  @Output()
  atOffsetMarker: EventEmitter<number> = new EventEmitter<number>();

  @ViewChild('textContainer')
  editableContent: ElementRef;

  destory$: EventEmitter<any> = new EventEmitter<any>();

  constructor(private renderer: Renderer2) {
  }

  ngOnInit(): void {
  }

  onKeypress() {
    this.textContentChange.next(this.editableContent.nativeElement.innerText);
    this.tryEmitOffset();
  }

  ngOnDestroy(): void {
    this.destory$.next(true);
    this.destory$.complete();
  }

  updateInnerHtml(value: string) {
    if (!value || !this.editableContent) {
      return;
    }
    this.editableContent.nativeElement.innerText = '';

    const lines = value.split('\n');

    lines.forEach((line: string) => {
      if (line.match(/^#OFFSET:.*/g)) {
        this.renderer.appendChild(this.editableContent.nativeElement, this.newOffsetElement(line));
      } else {
        const el = this.renderer.createElement('span');
        el.innerText = `${line}\n`;
        this.renderer.appendChild(this.editableContent.nativeElement, el);
      }
    });
  }

  private newOffsetElement(line: string): any {
    const el = this.renderer.createElement('span');
    el.innerText = `${line}\n`;
    el.className = 'do-not-edit';
    return el;
  }

  private newTextElement(text: string): any {
    const el = this.renderer.createElement('span');
    el.innerText = `${text}\n`;
    return el;
  }


  private tryEmitOffset() {
    const caretFocus = this.getCaretFocus();
    if (isOffsetLine(caretFocus.line)) {
      this.atOffsetMarker.next(getOffsetValueFromLine(caretFocus.line));
    }
  }

  insertOffsetAboveCaret(seconds: number): void {
    let sel = document.getSelection();
    let nd = sel.anchorNode;
    this.renderer.insertBefore(nd.parentNode, this.newOffsetElement(`#OFFSET: ${seconds}`), nd);
  }

  insertTextAboveCaret(text: string): void {
    let sel = document.getSelection();
    let nd = sel.anchorNode;
    this.renderer.insertBefore(nd.parentNode, this.newTextElement(text), nd);
  }

  getCaretFocus(): CaretFocus {
    let sel = document.getSelection();
    let nd = sel.anchorNode;
    return new CaretFocus(nd.textContent, sel.focusOffset);
  }

  ngAfterViewInit(): void {
    this.updateInnerHtml(this._textContent);
  }
}

class CaretFocus {
  constructor(public line: string, public offset: number) {
  }
}
